package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/berserktech/telebot/gh"
	"github.com/berserktech/telebot/tg"
)

// Handler
// =======

// IMPORTANT: the "println" calls in this function are mainly because I was
// struggling trying to set up the environment variables on Zeit.co
// Let's leave them where they are for now since we might continue playing around with the
// hosting platform. We can improve them, for sure.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Getting the message from GitHub
	secret := os.Getenv("GITHUB_CLIENT_SECRET")
	message, err := gh.GetMessage(r, secret)
	if err != nil {
		log.Print(err)
		fmt.Fprintf(w, "%s", err)
		return
	}
	println("Message:")
	println(message)

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		println("No token received")
	}

	// How to get the TELEGRAM_CHAT_ID: https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id
	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	println("Chat ID:", chatId)

	// Sending the message to Telegram
	if err := tg.Send(message, token, chatId); err != nil {
		log.Print(err)
		fmt.Fprintf(w, "%s", err)
		return
	}

	fmt.Fprintf(w, "Sent:\n%s", message)
}
