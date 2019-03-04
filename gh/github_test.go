package gh

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// NOTE:
// Take more payloads from: https://developer.github.com/v3/activity/events/types/

func eventRequest(event string, modifier string) *http.Request {
	path, _ := filepath.Abs(fmt.Sprintf("fixtures/github_%s%s.json", event, modifier))
	reader, _ := os.Open(path)

	request := httptest.NewRequest("POST", "/", reader)
	request.Header.Add("X-GitHub-Event", event)
	request.Header.Add("X-Hub-Signature", os.Getenv("GITHUB_SECRET"))
	return request
}

func TestGetMessageCommitComment(t *testing.T) {
	message, err := GetMessage(eventRequest("commit_comment", ""), "")
	assert.Nil(t, err)

	expected := "[Codertocat](https://github.com/Codertocat) commented one commit with:\n\nThis is a really good change! :+1:\n\nhttps://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860"
	assert.Equal(t, expected, message)
}

func TestGetMessageIssueComment(t *testing.T) {
	message, err := GetMessage(eventRequest("issue_comment", ""), "")
	assert.Nil(t, err)

	expected := "[Codertocat](https://github.com/Codertocat) commented one issue with:\n\nYou are totally right! I'll get this fixed right away.\n\nhttps://github.com/Codertocat/Hello-World/issues/2#issuecomment-393304133"
	assert.Equal(t, expected, message)
}

func TestGetMessagePullRequestReviewComment(t *testing.T) {
	message, err := GetMessage(eventRequest("pull_request_review_comment", ""), "")
	assert.Nil(t, err)

	expected := "[Codertocat](https://github.com/Codertocat) commented one pull request with:\n\nMaybe you should use more emojji on this line.\n\nhttps://github.com/Codertocat/Hello-World/pull/1#discussion_r191908831"
	assert.Equal(t, expected, message)
}

func TestGetMessagePullRequestReview(t *testing.T) {
	message, err := GetMessage(eventRequest("pull_request_review", ""), "")
	assert.Nil(t, err)

	expected := "[Codertocat](https://github.com/Codertocat) submitted the pull request review: Update the README with new information https://github.com/Codertocat/Hello-World/pull/1"
	assert.Equal(t, expected, message)
}

func TestGetMessagePullRequest(t *testing.T) {
	message, err := GetMessage(eventRequest("pull_request", ""), "")
	assert.Nil(t, err)

	expected := "[Codertocat](https://github.com/Codertocat) closed the pull request: Update the README with new information https://github.com/Codertocat/Hello-World/pull/1 Details:\nAdditions: 1 Deletions: 1"
	assert.Equal(t, expected, message)
}

func TestGetMessageIssues(t *testing.T) {
	message, err := GetMessage(eventRequest("issues", ""), "")
	assert.Nil(t, err)

	expected := "[Codertocat](https://github.com/Codertocat) opened the issue: Spelling error in the README file https://github.com/Codertocat/Hello-World/issues/2"
	assert.Equal(t, expected, message)
}

func TestGetMessageStatus(t *testing.T) {
	message, err := GetMessage(eventRequest("status", ""), "")
	assert.Nil(t, err)

	expected := "`success`: [Initial commit](https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240) by [Codertocat](https://github.com/Codertocat)"
	assert.Equal(t, expected, message)
}

func TestPing(t *testing.T) {
	message, err := GetMessage(eventRequest("ping", ""), "")
	assert.Nil(t, err)

	expected := "ping"
	assert.Equal(t, expected, message)
}

// Intentional failures:

func TestGetMessageStatusPending(t *testing.T) {
	_, err := GetMessage(eventRequest("status", "_pending"), "")
	assert.Equal(t, err, errors.New("gh: not allowed status, pending"))
}

func TestGetMessageIssuesLabeled(t *testing.T) {
	_, err := GetMessage(eventRequest("issues", "_edited"), "")
	assert.Equal(t, err, errors.New("gh: not allowed action, edited"))
}

func TestOrgBlockEventFailed(t *testing.T) {
	_, err := GetMessage(eventRequest("org_block", ""), "")
	assert.Equal(t, err, errors.New("event not defined to be parsed"))
}
