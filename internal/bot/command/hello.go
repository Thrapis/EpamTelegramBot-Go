package command

import (
	rq "baa-telebot/internal/bot/request"
	cfg "baa-telebot/internal/config"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func commandHello(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	config := cfg.GetConfig()
	result := rq.HelloRequest(config.Server)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	bot.Send(msg)
}
