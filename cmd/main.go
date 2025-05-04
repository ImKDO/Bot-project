package main

import (
	"Bot-project/internal/bot"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: %s", err)
	}

	token := os.Getenv("TG_API_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	tgBot := tgBot.NewBot(bot)
	tgBot.Start()
}
