package main

import (
	"net/smtp"
	"log"
	"crypto/tls"
	"net/mail"
	"net"
	"fmt"
	"os"
)

type message struct {
	title string
	body string
}

type mailService struct {
	email string
	username string
	password string
	server string
	port string
}

var service = &mailService{}

func init() {
	service.email = os.Getenv("TELETHINGS_EMAIL")
	service.username = os.Getenv("TELETHINGS_USER")
	service.password = os.Getenv("TELETHINGS_PASSWORD")
	service.server = os.Getenv("TELETHINGS_SERVER")
	service.port = os.Getenv("TELETHINGS_PORT")
}

//var botMail = os.Getenv("BOTMAIL")

func sendToThings(to string, m message) {

}



func ses() {

	from := mail.Address{"", service.email}
	to   := mail.Address{"", "add-to-things-ifbanxjdk0b70vyekl3@things.email"}
	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	//servername := service.server + ":" + service.port

	//host, _, _ := net.SplitHostPort(servername)


	auth := smtp.PlainAuth("", service.username, service.password, service.server)

	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: service.server,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", net.JoinHostPort(service.server, service.port), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, service.server)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}