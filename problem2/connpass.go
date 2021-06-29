package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func connpassMain() {
	endPoint := "https://connpass.com/api/v1/event/"
	resp, _ := http.Get(endPoint)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(XXXX)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error: ", err)
		return
	}
	fmt.Println(data.Events[0])
}
