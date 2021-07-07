package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const logFile = "tmp/logs.json" // 時間があったらsqliteに保存する
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

type WeatherLog struct {
	// Time    string `json:"time"`
	Weather string `json:"weather"`
	Temp    string `json:"temp"`
	Speed   string `json:"wind speed"`
}

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func main() {
	buildServer()
	t := time.NewTicker(360 * time.Second)
	for {
		select {
		case <-t.C:
			buildServer()
		}
	}
}
func buildServer() {
	println("server - http://localhost:8080")
	http.HandleFunc("/", showHandler)
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":8080", nil)
}

// func showWeatherData() {
// 	token := "4b729f5b5fb545d31c278041f43b99e2"
// 	city := "Tokyo,jp"
// 	endPoint := "https://api.openweathermap.org/data/2.5/weather"

// 	values := url.Values{}
// 	values.Set("q", city)
// 	values.Set("APPID", token)

// 	res, err := http.Get(endPoint + "?" + values.Encode())
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer res.Body.Close()
// 	bytes, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(string(bytes))
// 	// ここのbytesをopenweather.jsonに格納したい
// 	// json parse
// 	var apiRes OpenWeatherMapAPIResponse
// 	if err := json.Unmarshal(bytes, &apiRes); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(apiRes)
// 	weatherlog := weatherLog{
// 		// timeはその時間を出力
// 		Weather: apiRes.Weather[0].Main,
// 		Temp:    int(apiRes.Main.Temp - 273),
// 		Speed:   apiRes.Wind.Speed,
// 	}

// 	bytes, _ = json.Marshal(&weatherlog)
// 	ioutil.WriteFile(weatherFile, bytes, 0644)
// 	fmt.Println(string(bytes))
// }

func showHandler(w http.ResponseWriter, r *http.Request) {
	htmlLog := ""

	logs := loadLogs()
	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span>: %s --- %s</p>",
			i.ID,
			html.EscapeString(i.Name),
			html.EscapeString(i.Body),
			time.Unix(i.CTime, 0).Format("2006/1/2 15:04"))
	}

	token := "4b729f5b5fb545d31c278041f43b99e2"
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
	// fmt.Println(string(bytes))
	// ここのbytesをopenweather.jsonに格納したい
	// json parse
	var apiRes OpenWeatherMapAPIResponse
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
	fmt.Println(apiRes)
	// weatherlog := WeatherLog{
	// 	// timeはその時間を出力
	// 	Weather: apiRes.Weather[0].Main,
	// 	Temp:    string(int(apiRes.Main.Temp - 273)),
	// 	Speed:   string(int(apiRes.Wind.Speed)),
	// }

	// bytes, _ = json.Marshal(&weatherlog)
	// ioutil.WriteFile(weatherFile, bytes, 0644)
	// fmt.Println(string(bytes))
	// weatherLogAPI, err := ioutil.ReadFile(weatherFile)
	// if err != nil {
	// 	return
	// }

	// htmlLogAPI := string(weatherLogAPI)
	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>No71c3 8o4rD</h1><h2>Let's play with 1337 5p34k</h2>" +
		getForm() +
		// API
		"<p>Weather Broadcast at Tokyo<br>Weather : " + apiRes.Weather[0].Main + " (" + apiRes.Weather[0].Description + ")" + "<br>" +
		// "Wind speed : " + string(int(apiRes.Wind.Speed)) + "<br>" +
		"</p>" +
		"<p>Next Event in Connpass<br>" +
		// htmlLogAPI +
		"</p>" +
		htmlLog +
		"</body></html>"
	w.Write([]byte(htmlBody))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var log Log
	log.Name = r.Form["name"][0]
	log.Body = r.Form["body"][0]
	if log.Name == "" {
		log.Name = "4n0nym0u5"
	}
	logs := loadLogs()
	log.ID = len(logs) + 1
	log.CTime = time.Now().Unix()
	logs = append(logs, log)
	saveLogs(logs)
	http.Redirect(w, r, "/", 302)
}

func getForm() string {
	return "<div><form action='/write' method='POST'>" +
		"N4M3: <input type='text' name='name'><br>" +
		"CH47: <input type='text' name='body' style='width:30em;'><br>" +
		"<input type='submit' value='POST'>" +
		"</form></div><hr>"
}

func loadLogs() []Log {
	text, err := ioutil.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}
	var logs []Log
	json.Unmarshal([]byte(text), &logs)
	return logs
}

func saveLogs(logs []Log) {
	bytes, _ := json.Marshal(logs)
	ioutil.WriteFile(logFile, bytes, 0644)
}

// http://program.okitama.org/posts/2017-08-23_golang-timer-ticker/
// 3:04PM
