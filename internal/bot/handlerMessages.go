package tgBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleMessages(update tgbotapi.Update) {
	log.Printf("\n[%s] %s", update.Message.From.UserName, update.Message.Text, "\n")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommands(update.Message)
			continue
		}
		b.handleMessages(update)
	}
}
