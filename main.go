package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
) // Variables used for command line parameters

var (
	Token    string
	args     = make(map[string]func(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment))
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
	args["quote"] = handleQuote
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
