package main

import (
	"log"
	"net/http"
	"os"
	"net"
)

func main() {

	bot := setUpBot()
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), nil)
	runBot(bot)

}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm DndSpellsBot!"))
}