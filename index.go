package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/github"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

// IMPORTANT:
// I tried to separate this in several files, but Zeit didn't let me

// GitHub related code
// ===================

type Sender struct {
	Login string
}

type Comment struct {
	Body    string
	HTMLURL string
}

func parseComment(kind string, sender Sender, comment Comment) string {
	fmt.Println(kind)
	return fmt.Sprintf("*%s commented one %s:* %s %s", sender.Login, kind, comment.Body, comment.HTMLURL)
}

// Taken from: https://github.com/go-playground/webhooks/blob/v5/README.md
func getMessage(r *http.Request, secret string) (string, error) {
	// Handling the Github event
	hook, _ := github.New(github.Options.Secret(secret))
	payload, err := hook.Parse(r, github.CommitCommentEvent, github.IssueCommentEvent, github.PullRequestReviewCommentEvent)

	if err != nil {
		return "", err
	}

	switch payload.(type) {
	case github.CommitCommentPayload:
		p := payload.(github.CommitCommentPayload)
		return parseComment("commit", Sender{Login: p.Sender.Login}, Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}), nil

	case github.IssueCommentPayload:
		p := payload.(github.IssueCommentPayload)
		return parseComment("issue", Sender{Login: p.Sender.Login}, Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}), nil

	case github.PullRequestReviewCommentPayload:
		p := payload.(github.PullRequestReviewCommentPayload)
		return parseComment("pull request", Sender{Login: p.Sender.Login}, Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}), nil
	}

	return "", nil
}

// Telegram related code
// =====================

// Based on: https://github.com/go-telegram-bot-api/telegram-bot-api
func sendMessage(message string, token string, chatId string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	bot.Debug = true
	i64ID, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(i64ID, message)
	bot.Send(msg)
	return nil
}

// Handler
// =======

func Handler(w http.ResponseWriter, r *http.Request) {
	// Getting the message from GitHub
	secret := os.Getenv("GITHUB_SECRET")
	message, err := getMessage(r, secret)
	if err != nil {
		log.Panic(err)
	}

	token := os.Getenv("TELEGRAM_TOKEN")

	// How to get the TELEGRAM_CHAT_ID: https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	// Sending the message to Telegram
	if err := sendMessage(message, token, chatId); err != nil {
		log.Panic(err)
	}
}
