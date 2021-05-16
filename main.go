package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
) // Variables used for command line parameters

var (
	Token string
	args  = make(map[string]func(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment))
	animeUrl = "https://api.trace.moe/search?anilistInfo&url="
	)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not find Token")
		return
	}
	Token = os.Getenv("Token")
	args["ping"] = handlePing
	args["anime"] = handleImage
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	content := strings.Split(m.Content, " ")
	if content[0] == "" || content[0][0] != '!' {
		return
	}
	id := m.ChannelID
	attach := m.Attachments
	if val, ok := args[content[0][1:]]; ok {
		val(s, content[1:], id, attach)
	}
}

func handlePing(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment) {
	if len(m) == 0 {
		return
	}
	fmt.Println(len(m))
	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m[0]) == "ping" {
		fmt.Println("Pong!")
		s.ChannelMessageSend(id, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m[0]) == "pong" {
		fmt.Println("Ping!")
		s.ChannelMessageSend(id, "Ping!")
	}
}

func handleImage(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment) {
	if len(attach) == 0{
		return
	}
	resp, err := http.Get(animeUrl+attach[0].URL)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	var jResp AnimeResult
	if err:= json.NewDecoder(resp.Body).Decode(&jResp); err != nil {
		fmt.Println("something went wrong dude")
	}
	respMap := jResp.StringOutput(0)
	//fmt.Println(respMap["title_nat"])
	//fmt.Println(respMap["title_nat"])
	//fmt.Println(respMap["from"])
	//fmt.Println(respMap["video"])
	//fmt.Println(respMap["episode"])
	retStr := fmt.Sprintf("Title:%s\nEpisode:%s\nFrom:%s\nTo:%s",respMap["title_eng"],respMap["episode"],respMap["from"],respMap["to"])
	s.ChannelMessageSend(id,retStr)
	s.ChannelMessageSend(id, respMap["video"])
}
