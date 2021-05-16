package main

import (
	"fmt"
	"strconv"
)

type Anilist struct{
	Title Title `json:"title"`
}

type Title struct{
	Native string `json:"native"`
	English string `json:"english"`
}

type InResult struct{
	Anilist Anilist `json:"anilist"`
	Episode int `json:"episode"`
	From float64 `json:"from"`
	To float64 `json:"to"`
	Video string `json:"video"`
}

type AnimeResult struct {
	Result []InResult `json:"result"`
}

func (t AnimeResult) StringOutput(idx int) map[string]string {
	retMap := make(map[string]string)
	retMap["title_eng"] = t.Result[idx].Anilist.Title.English
	retMap["title_nat"] = t.Result[idx].Anilist.Title.Native
	retMap["episode"] = strconv.Itoa(t.Result[idx].Episode)
	retMap["from"] = fmt.Sprintf("%f", t.Result[idx].From)
	retMap["to"] = fmt.Sprintf("%f", t.Result[idx].To)
	retMap["video"] = t.Result[idx].Video
	return retMap
}