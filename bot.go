package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	//"os"
	//"net/http"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strings"
)

//var token = os.Getenv("TELETHINGS_BOT_TOKEN")
type config struct {
	Token string `json:"TELETHINGS_BOT_TOKEN"`
}

func setUpBot() *tgbotapi.BotAPI {
	log.Println("SETTING UP!")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.heroku.com/apps/mighty-wave-18558/config-vars", nil)
	if err != nil {
		log.Println("error while creating request instance")
		log.Panic(err)
	}
	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", "Basic a29uYXB0QGdtYWlsLmNvbTo4KXtrMUF5Pw==")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error getting response")
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read the body")
		log.Panic(err)
	}
	config := &config{}
	err = json.Unmarshal(body, config)
	if err != nil {
		log.Panic(err)
	}
	log.Println(config.Token)
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Println(config.Token)
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func runBot(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//updates, err := bot.GetUpdatesChan(u)

	_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://mighty-wave-18558.herokuapp.com/"+bot.Token, "server.crt"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	//go http.ListenAndServeTLS("0.0.0.0" + os.Getenv("PORT"), "server.crt", "server.key", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		handleCommands(update, bot)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func handleCommands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var resp, title, body string
	var msg = tgbotapi.NewMessage(update.Message.Chat.ID, resp)

	log.Println(update.Message.Command())
	switch update.Message.Command() {
	case "register":
	case "new":
		resp, title, body = getTitleBody(update.Message.Text)
		go sendToThings([]string{"konapt@gmail.com"}, title, body)
	case "delete":
	case "help":
		resp = `
Here is the list of coomands I understand:

/help     - this list of commands
/register - register your Things3 email
/delete   - remove your Things3 email from my database
/new      - add new note, syntax: '[title] | body'
`
	default:
		resp = "Please try using /help first"
		msg.ReplyToMessageID = update.Message.MessageID
	}
	msg.Text = resp
	bot.Send(msg)
}

func getTitleBody(text string) (string, string, string)  {
	var resp, title, body string
	resp = "Trying to send your note to Things3!"

	artifacts := strings.Split(text, "\t")
	switch len(artifacts) {
	case 0:
		resp = "Sorry, a note should have a title at least."
	case 1:
		title = artifacts[0]
		body = ""
	case 2:
		title = artifacts[0]
		body = artifacts[1]
	}
	return resp, title, body
}


func sendToThings(to []string, title, body string) {
	msg, err := json.Marshal(struct {
		To []string `json:"to"`
		Title string `json:"title"`
		Body string `json:"body"`
	}{to, title, body})
	if err != nil {
		log.Panicln("failed to marshall the message")
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://mailer-dot-telethings-196912.appspot.com/msg", bytes.NewReader(msg))
	if err != nil {
		log.Panicln("failed to create a new request")
	}
	client.Do(req)
}

func getUserThingsMail()  {

}