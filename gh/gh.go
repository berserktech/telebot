// Package gh handles github-related actions.
package gh

import (
	"fmt"
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

// Taken from: https://github.com/go-playground/webhooks/blob/v5/README.md
func GetMessage(r *http.Request, secret string) (string, error) {
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
		github.StatusEvent,
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
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		comment := Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}

		return comment.Format("commit", sender), nil

	case github.IssueCommentPayload:
		p := payload.(github.IssueCommentPayload)
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		comment := Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}

		return comment.Format("issue", sender), nil

	case github.PullRequestReviewCommentPayload:
		p := payload.(github.PullRequestReviewCommentPayload)
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		comment := Comment{Body: p.Comment.Body, HTMLURL: p.Comment.HTMLURL}

		return comment.Format("pull request", sender), nil

		// Events that have CRUD-like actions
	case github.PullRequestReviewPayload:
		p := payload.(github.PullRequestReviewPayload)
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		content := Content{Action: p.Action, Title: p.PullRequest.Title, HTMLURL: p.PullRequest.HTMLURL, Body: p.Review.Body}

		if err := content.NotAllowed(); err != nil {
			return "", err
		}

		return content.Format("pull request review", sender), nil

	case github.PullRequestPayload:
		p := payload.(github.PullRequestPayload)
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		body := fmt.Sprintf("Additions: %d Deletions: %d", p.PullRequest.Additions, p.PullRequest.Deletions)
		content := Content{Action: p.Action, Title: p.PullRequest.Title, HTMLURL: p.PullRequest.HTMLURL, Body: body}

		if err := content.NotAllowed(); err != nil {
			return "", err
		}

		return content.Format("pull request", sender), nil

	case github.IssuesPayload:
		p := payload.(github.IssuesPayload)
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		content := Content{Action: p.Action, Title: p.Issue.Title, HTMLURL: p.Issue.HTMLURL}

		if err := content.NotAllowed(); err != nil {
			return "", err
		}

		return content.Format("issue", sender), nil

		// Status are events triggered by commits
	case github.StatusPayload:
		p := payload.(github.StatusPayload)
		sender := Sender{Login: p.Sender.Login, HTMLURL: p.Sender.HTMLURL}
		status := Status{State: p.State, Message: p.Commit.Commit.Message, HTMLURL: p.Commit.HTMLURL}

		if err := status.NotAllowed(); err != nil {
			return "", err
		}

		return status.Format(sender), nil
		// Ping is simply so that we can run a minimal test.
	case github.PingPayload:
		return "ping", nil
	}

	return "", nil
}
