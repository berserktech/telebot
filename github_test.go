package main

import (
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

func eventRequest(event string) *http.Request {
	path, _ := filepath.Abs(fmt.Sprintf("fixtures/github_%s.json", event))
	reader, _ := os.Open(path)

	request := httptest.NewRequest("POST", "/", reader)
	request.Header.Add("X-GitHub-Event", event)
	return request
}

func TestGetMessageCommitComment(t *testing.T) {
	message, err := getMessage(eventRequest("commit_comment"), "")
	assert.Nil(t, err)

	expected := "*Codertocat commented one commit:* This is a really good change! :+1: https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860"
	assert.Equal(t, expected, message)
}

func TestGetIssueComment(t *testing.T) {
	message, err := getMessage(eventRequest("issue_comment"), "")
	assert.Nil(t, err)

	expected := "*Codertocat commented one issue:* You are totally right! I'll get this fixed right away. https://github.com/Codertocat/Hello-World/issues/2#issuecomment-393304133"
	assert.Equal(t, expected, message)
}

func TestGetPullRequestReviewComment(t *testing.T) {
	message, err := getMessage(eventRequest("pull_request_review_comment"), "")
	assert.Nil(t, err)

	expected := "*Codertocat commented one pull request:* Maybe you should use more emojji on this line. https://github.com/Codertocat/Hello-World/pull/1#discussion_r191908831"
	assert.Equal(t, expected, message)
}
