package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/go-playground/webhooks.v5/github"
	"log"
	"net/http"
	"os"
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

type Content struct {
	Action  string
	Title   string
	HTMLURL string
}

func parseComment(kind string, sender Sender, comment Comment) string {
	return fmt.Sprintf("%s commented one %s with:\n\n%s\n\n%s", sender.Login, kind, comment.Body, comment.HTMLURL)
}

func parseContent(kind string, sender Sender, content Content) string {
	return fmt.Sprintf("%s %s the %s: %s %s", sender.Login, content.Action, kind, content.Title, content.HTMLURL)
}

// Taken from: https://github.com/go-playground/webhooks/blob/v5/README.md
func getMessage(r *http.Request, secret string) (string, error) {
	// Handling the Github event
	hook, _ := github.New(github.Options.Secret(secret))
	payload, err := hook.Parse(r,
		// Comment events
		github.CommitCommentEvent,
		github.IssueCommentEvent,
		github.PullRequestReviewCommentEvent,
		// Events that are more relevant by their action
		github.PullRequestReviewEvent,
		github.PullRequestEvent,
		github.IssuesEvent,
		// Misc
		github.PingEvent)

	if err != nil {
		return "", err
	}

	switch payload.(type) {
	// Comment events
	case github.CommitCommentPayload:
		p := payload.(github.CommitCommentPayload)
		return parseComment("commit", Sender{Login: p.Sender.Login}, Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}), nil
	case github.IssueCommentPayload:
		p := payload.(github.IssueCommentPayload)
		return parseComment("issue", Sender{Login: p.Sender.Login}, Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}), nil
	case github.PullRequestReviewCommentPayload:
		p := payload.(github.PullRequestReviewCommentPayload)
		return parseComment("pull request", Sender{Login: p.Sender.Login}, Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}), nil

		// Events that are more relevant by their action
	case github.PullRequestReviewPayload:
		p := payload.(github.PullRequestReviewPayload)
		return parseContent("pull request review", Sender{Login: p.Sender.Login}, Content{Action: p.Action, Title: p.PullRequest.Title, HTMLURL: p.PullRequest.HTMLURL}), nil
	case github.PullRequestPayload:
		p := payload.(github.PullRequestPayload)
		return parseContent("pull request", Sender{Login: p.Sender.Login}, Content{Action: p.Action, Title: p.PullRequest.Title, HTMLURL: p.PullRequest.HTMLURL}), nil
	case github.IssuesPayload:
		p := payload.(github.IssuesPayload)
		return parseContent("issue", Sender{Login: p.Sender.Login}, Content{Action: p.Action, Title: p.Issue.Title, HTMLURL: p.Issue.HTMLURL}), nil

		// Misc
	case github.PingPayload:
		return "ping", nil
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
	// All group chat IDs are negative numbers, apparently
	msg := tgbotapi.NewMessage(-i64ID, message)
	bot.Send(msg)
	return nil
}

// Handler
// =======

func Handler(w http.ResponseWriter, r *http.Request) {
	// Getting the message from GitHub
	secret := os.Getenv("GITHUB_CLIENT_SECRET")
	message, err := getMessage(r, secret)
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
	if err := sendMessage(message, token, chatId); err != nil {
		log.Print(err)
		fmt.Fprintf(w, "%s", err)
		return
	}

	fmt.Fprintf(w, "Sent:\n%s", message)
}
