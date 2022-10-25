package command

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	cmdMessageHello = "hello"
	cmdMessageTime  = "time"
)

func RecognizeCommand(message string) func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if message == cmdMessageHello {
		return commandHello
	} else if message == cmdMessageTime {
		return commandTime
	}
	return nil
}
