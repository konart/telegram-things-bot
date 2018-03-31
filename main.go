package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	go http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), nil)
	go http.ListenAndServeTLS("0.0.0.0", "server.crt", "server.key", nil)
	go selfPing()
	bot := setUpBot()
	runBot(bot)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fmt.Fprintln(w, "Nothing to see here, human, move along.")
}

func selfPing() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://mighty-wave-18558.herokuapp.com/", nil)
	if err != nil {
		log.Panic("couldn't create request for selfPing")
	}
	for {
		time.Sleep(time.Minute * 30)
		client.Do(req)
	}
}
