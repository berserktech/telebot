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
// I tried to separate this in several files, but Zeit didn't let me.
// I'll continue investigating later.

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
	Body    string
}

// Builds up messages that follow a common pattern around a Comment struct.
// The messages will use a "kind" to identify the event in a humanly readable way,
// and two structs holding the data coming from the API, a Sender and a Comment.
// TODO: Perhaps use a template engine.
func parseComment(kind string, sender Sender, comment Comment) string {
	return fmt.Sprintf("*%s* commented one %s with:\n\n%s\n\n%s", sender.Login, kind, comment.Body, comment.HTMLURL)
}

// Builds up messages that have CRUD-like actions
// The messages will use a "kind" to identify the event in a humanly readable way,
// and two structs holding the data coming from the API, a Sender and a Content.
// The output vries if the provided Content has a Body.
// TODO: Perhaps use a template engine.
func parseCRUD(kind string, sender Sender, content Content) string {
	var body string
	if content.Body != "" {
		body = fmt.Sprintf(" Details:\n%s", content.Body)
	}
	return fmt.Sprintf("*%s* %s the %s: %s %s%s", sender.Login, content.Action, kind, content.Title, content.HTMLURL, body)
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
		// Events that have CRUD-like actions
		github.PullRequestReviewEvent,
		github.PullRequestEvent,
		github.IssuesEvent,
		// Misc
		github.PingEvent)

	if err != nil {
		return "", err
	}

	// NOTES:
	// - The cases can't fallthrough when they belong to a switch over types.
	// - I'm trying to pass objects of a well defined struct to make the parsing functions smaller,
	//   since this switch is pretty verbose anyway.

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

		// Events that have CRUD-like actions
	case github.PullRequestReviewPayload:
		p := payload.(github.PullRequestReviewPayload)
		return parseCRUD("pull request review", Sender{Login: p.Sender.Login}, Content{Action: p.Action, Title: p.PullRequest.Title, HTMLURL: p.PullRequest.HTMLURL, Body: p.Review.Body}), nil
	case github.PullRequestPayload:
		p := payload.(github.PullRequestPayload)
		body := fmt.Sprintf("Additions: %d Deletions: %d", p.PullRequest.Additions, p.PullRequest.Deletions)
		return parseCRUD("pull request", Sender{Login: p.Sender.Login}, Content{Action: p.Action, Title: p.PullRequest.Title, HTMLURL: p.PullRequest.HTMLURL, Body: body}), nil
	case github.IssuesPayload:
		p := payload.(github.IssuesPayload)
		return parseCRUD("issue", Sender{Login: p.Sender.Login}, Content{Action: p.Action, Title: p.Issue.Title, HTMLURL: p.Issue.HTMLURL}), nil

		// Misc
		// Ping is simply so that we can run a minimal test.
	case github.PingPayload:
		return "ping", nil
	}

	return "", nil
}

// Telegram related code
// =====================

// Based on: https://github.com/go-telegram-bot-api/telegram-bot-api
// TODO: The configuration we set here is probably better in a configuration file.
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
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true
	bot.Send(msg)
	return nil
}

// Handler
// =======

// IMPORTANT: the "println" calls in this function are mainly because I was
// struggling trying to set up the environment variables on Zeit.co
// Let's leave them where they are for now since we might continue playing around with the
// hosting platform. We can improve them, for sure.
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
