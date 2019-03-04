package tg

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Based on: https://github.com/go-telegram-bot-api/telegram-bot-api
// TODO: The configuration we set here is probably better in a configuration file.
func Send(message string, token string, chatId string) error {
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
