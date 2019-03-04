# telebot 🤖✨

Telegram Bot for our GitHub community 🤗

![](https://user-images.githubusercontent.com/417016/53699784-eca52200-3db9-11e9-9477-fd6dab0ed4d5.png)

## Index

- [What is telebot](#what-is-telebot)
- [Why?](#why)
- [Supported Events](#supported-events)
- [How to build](#how-to-build)
- [How to contribute](#how-to-contribute)
- [How to deploy](#how-to-deploy)
- [License](#license)
- [References](#references)


## What is telebot

_telebot_ is a bot that sends each one of the [GitHub Webhooks][GHWH] we care
about to our private Telegram group.

Besides that, it's an example of a Go project that uses [GitHub
Webhooks][GHWH] and the [Telegram
API](https://github.com/go-telegram-bot-api/telegram-bot-api).

## Why?

This started as a go practice by [@sadasant][sadasant], but now it is
the official Telegram bot for BerserkTech.

My study notes can be found:
[here](https://github.com/berserktech/ideas/blob/master/estudios/Go/2019-03/after_3_years_without_go_en.md).

## Supported events

This bot is currently listening to the following webhook events (keep
in mind that some line breaks won't appear in this table):

| Event Name  | Output |
| ------------- | ------------- |
| [commit_comment](https://developer.github.com/v3/activity/events/types/#commitcommentevent) | [Codertocat](https://github.com/Codertocat) commented one commit with: This is a really good change!  :+1: https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860 |
| [issue_comment](https://developer.github.com/v3/activity/events/types/#issuecommentevent) | [Codertocat](https://github.com/Codertocat) commented one issue with: You are totally right! I'll get this fixed right away. https://github.com/Codertocat/Hello-World/issues/2#issuecomment-393304133 |
| [pull_request_review_comment](https://developer.github.com/v3/activity/events/types/#pullrequestreviewcommentevent) | [Codertocat](https://github.com/Codertocat) commented one pull request with: Maybe you should use more emojji on this line. https://github.com/Codertocat/Hello-World/pull/1#discussion_r191908831 |
| [pull_request_review](https://developer.github.com/v3/activity/events/types/#pullrequestreviewevent) | [Codertocat](https://github.com/Codertocat) submitted the pull request review: Update the README with new information https://github.com/Codertocat/Hello-World/pull/1 |
| [pull_request](https://developer.github.com/v3/activity/events/types/#pullrequestevent) | [Codertocat](https://github.com/Codertocat) closed the pull request: Update the README with new information https://github.com/Codertocat/Hello-World/pull/1 Details: ditions: 1 Deletions: 1 |
| [issues](https://developer.github.com/v3/activity/events/types/#issuesevent) | [Codertocat](https://github.com/Codertocat) edited the issue: Spelling error in the README file https://github.com/Codertocat/Hello-World/issues/2 |
| [status](https://developer.github.com/v3/activity/events/types/#statusevent) | `success`: [Initial commit](https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240) by [Codertocat](https://github.com/Codertocat) |
| [ping](https://developer.github.com/webhooks/#ping-event) | ping |

We should definitely add more and improve what we're currently doing
with each one of these events (check out the open issues!).

Some of the events are filtered. In detail:

- `status` if they have state equal to `pending`.
- `issues` if they have action equal to `labeled`, `unlabeled`,
  `assigned`, `unassigned`, `review_requested`,
  `review_request_removed` or `edited`.

## How to build

### Install Go

Make sure you have Go installed: <https://golang.org/doc/install>.
You can also use @[stefanmaric](https://github.com/stefanmaric)'s
[Simple go version manager, gluten-free](https://github.com/stefanmaric/g) 🙌

## How to contribute

Make an issue or a pull request! :) Remember to `fmt` your `.go`s 😆

There are also some useful commands at the root of this directory.
They're bash executables since that's how [Go's
source](https://github.com/golang/go/tree/master/src) handles similar
scripts. The list follows:

- `test.bash`: Runs `go test` in all the submodules of this repo.
- `fmt.bash`: Runs `go fmt` in all the submodules of this repo.
- `fmt-check.bash`: Runs `go fmt -l .` in all the submodules of this
  repo. If it finds files that don't have the proper formatting, it
  will exit with status code 1. We use this one for CI purposes
- `mod-update.bash`: Updates all the versions of local dependencies of
  this module to the latest hash in your local git. Make sure to have
  pushed it. To reconcile with Go's authomatic changes, it runs
  `fmt.bash` at the end, which will also check if the version is
  available through the network. To use this, you must pass a
  git-reference to update to. You can use `HEAD`, as in the following
  example: `./mod-update.bash HEAD~1`.

## How to deploy

### Creating a Telegram Bot

Go to <https://telegram.me/botfather> and follow the steps :)

Make sure to store the HTTP API token in a safe place!

### Cloning this repository

Make sure you have git installed. Follow the GitHub guides:
<https://help.github.com/en#dotcom>, they're way better than anything
I can come up with.

Once you have `git`, you can clone this repo with:

```
git clone https://github.com/berserktech/telebot
```

or:

```
git clone git@github.com:berserktech/telebot.git
```

### Install Zeit's Now

Install Zeit's `now` by going to: <https://zeit.co/download#now-cli>,
or running `npm install -g now`.
 
### Setting up the secrets

After you have `now`, you'll need to add the following secrets:

- github-secret: Your Webhook secret (more below).
- telegram-chat-id: The ID of your Telegram chat. `telebot` doesn't
  listen to telegram incoming messages, so you will need to follow the
  steps described here: <https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id>
- telegram-token: The HTTP API token obtained from the creation of the
  Telegram bot.

At the end, if you run `now secret ls`, it should look like this:

```
now secret ls
> 4 secrets found under sadasant [345ms]

  name                  created
  github-secret         9h ago
  telegram-chat-id      8h ago
  telegram-token        8h ago
```

### Deploy this project

As long as you have this project locally, you can run `now` at the
root of `telebot` to deploy it in with Zeit 👍 If you run it, you
should eventually get this output:

```
now
> Deploying ~/code/github.com/berserktech/telebot under sadasant
> Using project telebot
> Synced 2 files (6.91KB) [1s]
> https://telebot-[something random].now.sh [v2] [3s]
┌ index.go        Ready               [42s]
└── λ index.go (4.68MB) [iad1]
> Success! Deployment ready [45s]
```

At this point, you should be able to make network requests against it,
or to see the logs: `now logs https://telebot-[something random].now.sh`.

### Make a new GitHub Webhook Application

- Go to: <https://github.com/settings/apps>
- Click on `New GitHub App`.
- Fill at least the following fields in the form:
    - GitHub App name: Your app name. It can be anything.
    - User authorization callback URL: This is required, but we won't
      be using it, so put any valid URL here.
    - Webhook URL: Paste the URL generated by Zeit's Now. It should
      look like: `https://telebot-[something random].now.sh`.
    - Webhook secret: Set the same value you created in the
      `github-secret` secret.
    - Add as many permissions as you want. Keep in mind that you
      should probably allow this application to have read access to:
      commits, issues and pull requests.
    - Subscribe to all the events that you want. Keep in mind that
      telebot only answers to the events listed in: [Supported events](#supported-events).
- Save.
- At this point there should be a button list at the left, with a
  specific option: Install App. Go there and install your freshly
  created application to your account or organization 🙌 You're done!
 
## License

MIT, check the [LICENSE](/LICENSE) file.

## References
- [@sadasant's study notes](https://github.com/berserktech/ideas/blob/master/estudios/Go/2019-03/after_3_years_without_go_en.md)
- [Enabling module support for Go v1.11 in CircleCI](https://circleci.com/blog/go-v1.11-modules-and-circleci/)
- [GitHub REST API v3 - Event Types & Payloads](https://developer.github.com/v3/activity/events/types/)
- [Library webhooks allows for easy receiving and parsing of GitHub,
  Bitbucket and GitLab Webhook
  Events](https://github.com/go-playground/webhooks)
- [Telegram Bot API](https://core.telegram.org/bots/api)
- [Zeit by example | Go](https://zeit.co/examples/go)
- [Golang bindings for the Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)

[GHWH]: https://developer.github.com/webhooks/
[sadasant]: https://github.com/sadasant
