package main

import (
	cfg "baa-telebot/config"
	cl "baa-telebot/internal/bot/client"
	"fmt"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	cmdHello = "hello"
)

func main() {
	config := cfg.GetConfig()
	fmt.Println("Bot started...")
	StartBot(config.Token)
}

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if IsTextMessage(&update.Message.Text) {
			switch update.Message.Text {
			case cmdHello:
				CommandHello(bot, &update)

			}
		}
	}
}

func IsTextMessage(text *string) bool {
	return reflect.TypeOf(*text).Kind() == reflect.String && *text != ""
}

func CommandHello(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	config := cfg.GetConfig()
	result := cl.HelloRequest(config.Server)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	bot.Send(msg)
}
