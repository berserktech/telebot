# telebot
Telegram Bot for our repository

## Why?

This started as a go practice by
[@sadasant](https://github.com/sadasant), but now it is the official
Telegram bot for BerserkTech.

It doesn't do much, it listens to github events and sends them over
Telegram.


## Events

This bot is currently listening to the following events:


```
commit_comment
issue_comment
pull_request_review_comment
pull_request_review
pull_request
issues
status
ping
```

Plase make sure to check the unit tests we have so far.

## TODOs

- Move it to the BerserkTech organization in Zeit.
- Improve this readme.
- Write better integration tests.
- Get feedback.

## Useful links:
- [Enabling module support for Go v1.11 in CircleCI](https://circleci.com/blog/go-v1.11-modules-and-circleci/)
- [GitHub REST API v3 - Event Types & Payloads](https://developer.github.com/v3/activity/events/types/)
- [Library webhooks allows for easy receiving and parsing of GitHub,
  Bitbucket and GitLab Webhook
  Events](https://github.com/go-playground/webhooks)
- [Telegram Bot API](https://core.telegram.org/bots/api)
- [Zeit by example | Go](https://zeit.co/examples/go)
- [Golang bindings for the Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)
