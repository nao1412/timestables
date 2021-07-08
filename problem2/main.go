package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const logFile = "tmp/logs.json" // 時間があったらsqliteに保存する

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
	Weather string `json:"weather"`
	Temp    string `json:"temp"`
	Speed   string `json:"wind speed"`
}

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

func showHandler(w http.ResponseWriter, r *http.Request) {
	htmlLog := ""

	logs := loadLogs()
	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span> : %s --- %s</p>",
			i.ID,
			html.EscapeString(i.Name),
			html.EscapeString(i.Body),
			time.Unix(i.CTime, 0).Format("2006/1/2 15:04"))
	}

	weatherToken := "4b729f5b5fb545d31c278041f43b99e2"
	city := "Tokyo,jp"
	weatherEndpoint := "https://api.openweathermap.org/data/2.5/weather"

	values := url.Values{}
	values.Set("q", city)
	values.Set("APPID", weatherToken)

	res, err := http.Get(weatherEndpoint + "?" + values.Encode())
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
	// fmt.Println("openeweatherAPI", apiRes)

	day := time.Now()
	const layout = "20060102"
	const layout2 = "2006-01-02 15:04:05"
	const layout3 = "Mon, January 2, 2006"
	// fmt.Println(day.Format(layout))
	today := day.Format(layout)
	nowTime := day.Format(layout2)
	todayEvent := day.Format(layout3)
	connpassEndpoint := "https://connpass.com/api/v1/event/?ymd="
	connpassEndpoint += today
	resp, err := http.Get(connpassEndpoint)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)
	data := new(ConnpassEvent)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error: ", err)
		return
	}
	connpassEventTitle1 := data.Events[0].Title
	connpassEventURL1 := "<a href=\"" + data.Events[0].EventURL + "\">"
	connpassEventTitle2 := data.Events[1].Title
	connpassEventURL2 := "<a href=\"" + data.Events[1].EventURL + "\">"
	connpassEventTitle3 := data.Events[2].Title
	connpassEventURL3 := "<a href=\"" + data.Events[2].EventURL + "\">"
	connpassEventTitle4 := data.Events[3].Title
	connpassEventURL4 := "<a href=\"" + data.Events[3].EventURL + "\">"
	connpassEventTitle5 := data.Events[4].Title
	connpassEventURL5 := "<a href=\"" + data.Events[4].EventURL + "\">"

	leeturl := "https://en.wikipedia.org/wiki/Leet"
	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>No71c3 8o4rD 1337</h1><h2>Let's play with 1337 5p34k</h2>" +
		"<p>how to write leet speaks -> <a href=" + leeturl + ">click here</a></p>" +
		getForm() +
		// API
		"<p>Weather Broadcast at Tokyo (" + nowTime + ")<br>Weather : " + apiRes.Weather[0].Main + " (" + apiRes.Weather[0].Description + ")" + "<br>" +
		"Temperature : " + strconv.Itoa(int(apiRes.Main.Temp-273)) + " ℃<br>" +
		"Wind Speed : " + strconv.FormatFloat(apiRes.Wind.Speed, 'f', 2, 64) + " m/s<br>" +
		"</p>" +
		"<p>Today's Event in Connpass (" + todayEvent + ")<br>" +
		"1 : " + connpassEventURL1 + connpassEventTitle1 + "</a>" + "<br>" +
		"2 : " + connpassEventURL2 + connpassEventTitle2 + "</a>" + "<br>" +
		"3 : " + connpassEventURL3 + connpassEventTitle3 + "</a>" + "<br>" +
		"4 : " + connpassEventURL4 + connpassEventTitle4 + "</a>" + "<br>" +
		"5 : " + connpassEventURL5 + connpassEventTitle5 + "</a>" + "<br>" +
		"</p>" +
		"<h2>Comments</h2>" +
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
		"N4M3 : <input type='text' name='name'><br>" +
		"CH47 : <input type='text' name='body' style='width:30em;'><br>" +
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
