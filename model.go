package main

import (
	"fmt"
	"math/rand"
	"strconv"
)


//Types of titles
type Title struct{
	Native string `json:"native"`
	English string `json:"english"`
}

//Title of the anime
type Titles struct{
	Title Title `json:"title"`
}

//Values inside of the result key
type InResult struct{
	Titles Titles `json:"anilist"`
	Episode int `json:"episode"`
	From float64 `json:"from"`
	To float64 `json:"to"`
	Video string `json:"video"`
}

//List of results from json
type AnimeResult struct {
	Result []InResult `json:"result"`
}

type AnimeResultAlt struct {
	Result []InResultAlt `json:"result"`
}

//Optional result with episode list
type InResultAlt struct{
	Titles Titles `json:"anilist"`
	Episode []int `json:"episode"`
	From float64 `json:"from"`
	To float64 `json:"to"`
	Video string `json:"video"`
}

//Returns a map of the values in the response
func (t AnimeResult) MapOutput(idx int) map[string]string {
	retMap := make(map[string]string)
	retMap["title_eng"] = t.Result[idx].Titles.Title.English
	retMap["title_nat"] = t.Result[idx].Titles.Title.Native
	retMap["episode"] = strconv.Itoa(t.Result[idx].Episode)
	retMap["from"] = fmt.Sprintf("%.2f", t.Result[idx].From/60)
	retMap["to"] = fmt.Sprintf("%.2f", t.Result[idx].To/60)
	retMap["video"] = t.Result[idx].Video
	return retMap
}

func (t AnimeResultAlt) MapOutput(idx int) map[string]string {
	retMap := make(map[string]string)
	retMap["title_eng"] = t.Result[idx].Titles.Title.English
	retMap["title_nat"] = t.Result[idx].Titles.Title.Native
	retMap["episode"] = strconv.Itoa(t.Result[idx].Episode[0])
	retMap["from"] = fmt.Sprintf("%.2f", t.Result[idx].From/60)
	retMap["to"] = fmt.Sprintf("%.2f", t.Result[idx].To/60)
	retMap["video"] = t.Result[idx].Video
	return retMap
}

//Individual quote from
type Result struct{
	Anime string `json:"anime"`
	Character string `json:"character"`
	Quote string `json:"quote"`
}
//List of quotes from the result
type Quotes struct{
	Result []Result
}

//Returns a random quote from the character
func (t Quotes) MapOutput() map[string]string{
	idx := rand.Intn(5)
	retMap := t.Result[idx].MapOutput()
	return retMap
}

//Returns a map from the quote
func (t Result) MapOutput() map[string]string{
	retMap := make(map[string]string)
	retMap["anime"]=t.Anime
	retMap["character"] = t.Character
	retMap["quote"] = t.Quote
	return retMap
}