package main

import (
	"log"
	"net/http"
	"os"
	"net"
)

func main() {

	bot := setUpBot()
	runBot(bot)

	log.Fatal(http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), nil))

}
