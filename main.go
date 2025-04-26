package main

import (
	tgBot "Bot-project/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7976782538:AAFZ7QH1g5Z-UZxx6Za-9gonX3yHcncFgXg")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	tgBot := tgBot.NewBot(bot)
	tgBot.Start()
}
