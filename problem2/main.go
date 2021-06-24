package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"time"
)

const logFile = "logs.json"

type Log struct {
	ID    int    `json: "id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func main() {
	println("sever - http://localhost:8888")
	http.HandleFunc("/", showHandler)
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":8888", nil)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	htmlLog := ""
	logs := loadLogs()
	for _, i := range logs {
		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span>: %s --- %s</p>",
			i.ID,
			html.EscapeString(i.Name),
			html.EscapeString(i.Body),
			time.Unix(i.CTime, 0).Format("2006/1/1 15:04"))
	}
	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>にの</h1>" +
		getForm() + htmlLog + "</body></html>"
	w.Write([]byte(htmlBody))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var log Log
	log.Name = r.Form["name"][0]
	log.Body = r.Form["body"][0]
	if log.Name == "" {
		log.Name = "匿名希望"
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
		"なめえ: <input type='text' name='name'><br>" +
		"本文: <input type='text' name='body' style='width:30em;'><br>" +
		"<input type='submit' value='かきこみ'>" +
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
