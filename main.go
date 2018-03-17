package main

import (
	"net/http"
	"fmt"
	"net"
	"os"
	"log"
	"time"
)

func main() {
	go selfPing()
	http.HandleFunc("/", indexHandler)
	go http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), nil)
	go http.ListenAndServeTLS("0.0.0.0","server.crt", "server.key", nil)
	bot := setUpBot()
	runBot(bot)


}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
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
		time.Sleep(time.Second*15)
		client.Do(req)
	}
}