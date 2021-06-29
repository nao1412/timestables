package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const weatherFile = "openweather.txt"

type OpenWeatherMapAPIResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int64 `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

// type weatherLog struct {
// 	// Time    string `json:"time"`
// 	Weather string  `json:"weather"`
// 	Temp    int `json:"temp"`
// }

func saveWeatherLog(logs []weatherLog) {
	bytes, _ := json.Marshal(logs)
	ioutil.WriteFile(weatherFile, bytes, 0644) // main.goと同じ0644でも大丈夫？
}

func maain() {
	token := "XXXX"
	city := "Tokyo,jp"
	endPoint := "https://api.openweathermap.org/data/2.5/weather"

	values := url.Values{}
	values.Set("q", city)
	values.Set("APPID", token)

	res, err := http.Get(endPoint + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	// ここのbytesをopenweather.jsonに格納したい
	// json parse
	var apiRes OpenWeatherMapAPIResponse
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}

	weatherlog := weatherLog{
		// timeはその時間を出力
		Weather: apiRes.Weather[0].Description,
		Temp:    int(apiRes.Main.Temp - 273),
	}

	bytes, _ = json.Marshal(&weatherlog)
	ioutil.WriteFile(weatherFile, bytes, 0644)
	fmt.Println(string(bytes))

	fmt.Println(apiRes)

	fmt.Printf("時刻:%s\n", time.Unix(apiRes.Dt, 0))
	fmt.Printf("天気:%s\n", apiRes.Weather[0].Main)
	fmt.Printf("アイコン: https://openweathermap.org/img/wn/%s@2x.png\n", apiRes.Weather[0].Icon)
	fmt.Printf("説明:%s\n", apiRes.Weather[0].Description)
	fmt.Printf("緯度:%f°\n", apiRes.Coord.Lat)
	fmt.Printf("経度:%f°\n", apiRes.Coord.Lon)
	fmt.Printf("気温:%f℃\n", apiRes.Main.Temp-273)
	fmt.Printf("最高気温:%f℃\n", apiRes.Main.TempMax-273)
	fmt.Printf("最低気温:%f℃\n", apiRes.Main.TempMin-273)
	fmt.Printf("気圧: %dPa\n", apiRes.Main.Pressure)
	fmt.Printf("湿度: %d\n", apiRes.Main.Humidity)
	fmt.Printf("風速:%fm/s\n", apiRes.Wind.Speed)
	fmt.Printf("風向き:%d°\n", apiRes.Wind.Deg)

	weatherLogAPI, err := ioutil.ReadFile("openweather.txt")
	fmt.Println(string(weatherLogAPI))
	// """
	// var t string = apiRes.Weather[0].Description
	// weathertxt := "時刻:" + t + "\n"

	// ioutil.WriteFile(weatherFile, []byte(weathertxt), 0644)
	// """
}

// apiResをopenweather.jsonに保存して1時間ごとに変える->txtでもいいかな
