package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ConnpassEvent struct {
	ResultsStart     int `json:"results_start"`
	ResultsReturned  int `json:"results_returned"`
	ResultsAvailable int `json:"results_available"`
	Events           []struct {
		EventID          int         `json:"event_id"`
		Title            string      `json:"title"`
		Catch            string      `json:"catch"`
		Description      string      `json:"description"`
		EventURL         string      `json:"event_url"`
		StartedAt        time.Time   `json:"started_at"`
		EndedAt          time.Time   `json:"ended_at"`
		Limit            int         `json:"limit"`
		HashTag          string      `json:"hash_tag"`
		EventType        string      `json:"event_type"`
		Accepted         int         `json:"accepted"`
		Waiting          int         `json:"waiting"`
		UpdatedAt        time.Time   `json:"updated_at"`
		OwnerID          int         `json:"owner_id"`
		OwnerNickname    string      `json:"owner_nickname"`
		OwnerDisplayName string      `json:"owner_display_name"`
		Place            string      `json:"place"`
		Address          string      `json:"address"`
		Lat              interface{} `json:"lat"`
		Lon              interface{} `json:"lon"`
		Series           struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"series"`
	} `json:"events"`
}

func main() {
	endPoint := "https://connpass.com/api/v1/event/?keyword=elixir&count=1"
	resp, _ := http.Get(endPoint)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	jsonBytes := ([]byte)(byteArray)
	data := new(ConnpassEvent)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error: ", err)
		return
	}
	fmt.Println(data.Events[0])
}
