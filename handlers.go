package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
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
	s.ChannelMessageSend(id,"Quote")
}
