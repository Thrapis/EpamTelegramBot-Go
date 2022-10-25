package main

import (
	cmd "baa-telebot/internal/bot/command"
	cfg "baa-telebot/internal/config"

	"fmt"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func IsTextMessage(text *string) bool {
	return reflect.TypeOf(*text).Kind() == reflect.String && *text != ""
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
			if command := cmd.RecognizeCommand(update.Message.Text); command != nil {
				command(bot, &update)
			}
		}
	}
}

func main() {
	config := cfg.GetConfig()
	fmt.Println("Bot started...")
	StartBot(config.Token)
}
