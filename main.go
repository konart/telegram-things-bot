package main

import (
	"net/http"
	"fmt"
	//"net"
	"os"
)

func main() {

	http.HandleFunc("/", indexHandler)
	//go http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), nil)
	go http.ListenAndServeTLS("0.0.0.0" + os.Getenv("PORT"), "server.crt", "server.key", nil)
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