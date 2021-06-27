package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const weatherFile = "openweather.json"

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

func main() {
	token := "xxxxxx"
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
	fmt.Println(res, 1)
	fmt.Println(res.Body, 2)
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes, 3)
	fmt.Println(string(bytes), 4)
	// ここのbytesをopenweather.jsonに格納したい
	// saveWeatherLog(string(bytes))

	// json parse
	var apiRes OpenWeatherMapAPIResponse
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
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
}

func saveWeatherLog(logs []OpenWeatherMapAPIResponse) {
	bytes, _ := json.Marshal(logs)
	ioutil.WriteFile(weatherFile, bytes, 0644) // main.goと同じ0644でも大丈夫？
}
