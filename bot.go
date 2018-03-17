package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	//"os"
	//"net/http"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//var token = os.Getenv("TELETHINGS_BOT_TOKEN")
type config struct {
	Token string `json:"TELETHINGS_BOT_TOKEN"`
}

func setUpBot() *tgbotapi.BotAPI {
	log.Println("SETTING UP!")
	resp, err := http.Get("https://api.heroku.com/apps/mighty-wave-18558/config-vars")
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

		msg := handleCommands(update)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		bot.Send(msg)
	}
}

func handleCommands(update tgbotapi.Update) tgbotapi.MessageConfig {
	var resp string
	var msg = tgbotapi.NewMessage(update.Message.Chat.ID, resp)

	log.Println(update.Message.Command())
	switch update.Message.Command() {
	case "register":
	case "new":
		sendToThings([]string{"konapt@gmail.com"}, "Testing delivery", "Test successful")
		resp = "Trying to send your note to Things3!"
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
	return msg
}


func sendToThings(to []string, title, body string) {

}