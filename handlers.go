package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

//Gets called whenever a message is received
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	content := strings.Split(m.Content, " ")
	//If it is empty or not a command, ignore
	if content[0] == "" || content[0][0] != '!' {
		return
	}
	//if it is a valid command
	if val, ok := args[content[0][1:]]; ok {
		id := m.ChannelID
		attach := m.Attachments
		val(s, content[1:], id, attach)
	}
}

//!ping command
//Returns Ping! or Pong! to the user accordingly
func handlePing(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment) {
	//If there is nothing after the ping command
	if len(m) == 0 {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m[0]) == "ping" {
		s.ChannelMessageSend(id, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m[0]) == "pong" {
		s.ChannelMessageSend(id, "Ping!")
	}
}

//!anime command
//Calls trace.moe API on an image attachment and returns info on the image
func handleImage(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment) {
	//If no image is attached
	if len(attach) == 0{
		return
	}
	//Call API
	resp, err := http.Get(animeUrl+attach[0].URL)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	var jResp AnimeResult
	//Parse response into a go struct
	if err:= json.NewDecoder(resp.Body).Decode(&jResp); err != nil {
		fmt.Println("something went wrong dude")
	}
	//Returns the struct as a map
	respMap := jResp.MapOutput(0)
	//String to be returned to the user
	retStr := fmt.Sprintf("Title: %s\nEpisode: %s\nFrom: %s\nTo: %s",respMap["title_eng"],respMap["episode"],respMap["from"],respMap["to"])
	s.ChannelMessageSend(id,retStr)
	s.ChannelMessageSend(id, respMap["video"])
}

//!quote command
//Returns a quote based on the user's query
func handleQuote(s *discordgo.Session, m []string, id string, attach []*discordgo.MessageAttachment){
	newUrl := quotesUrl
	retStr := ""
	if len(m) == 0{
		newUrl=newUrl+"random"
		resp, err := http.Get(newUrl)
		if err != nil{
			fmt.Println("Something went wrong")
			return
		}
		defer resp.Body.Close()
		var randQuote Result
		bod, err := ioutil.ReadAll(resp.Body)
		if err:= json.Unmarshal(bod,&randQuote); err != nil {
			fmt.Println("something went wrong dude")
		}
		fmt.Println(randQuote)
		respMap := randQuote.MapOutput()
		retStr = fmt.Sprintf("Anime: %s\nCharacter: %s\nQuote: %s\n",respMap["anime"],respMap["character"],respMap["quote"])
	} else{
		page := rand.Intn(3)
		newUrl = newUrl+"quotes/character?name="+strings.Join(m," ")+"&page="+strconv.Itoa(page)
		fmt.Println(newUrl)
		resp, err := http.Get(newUrl)
		if err != nil{
			fmt.Println("Something went wrong")
			return
		}
		defer resp.Body.Close()
		bod, err := ioutil.ReadAll(resp.Body)
		var randQuote []Result
		if err:= json.Unmarshal(bod,&randQuote); err != nil {
			fmt.Println(err)
			fmt.Println("something went wrong dude")
		}
		idx := rand.Intn(5)
		respMap := randQuote[idx].MapOutput()
		fmt.Println(randQuote)
		//respMap := randQuote.MapOutput()
		retStr = fmt.Sprintf("Anime: %s\nCharacter: %s\nQuote: %s\n",respMap["anime"],respMap["character"],respMap["quote"])
	}
	s.ChannelMessageSend(id,retStr)
}
