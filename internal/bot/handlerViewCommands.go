package tgBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) setCommands(commands map[string]string) []tgbotapi.BotCommand {
	var setCommands []tgbotapi.BotCommand
	for c, d := range commands {
		setCommands = append(setCommands, tgbotapi.BotCommand{Command: c, Description: d})
	}
	return setCommands
}

func (b *Bot) handlerShowDownList() error {
	commands := b.setCommands(description)
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	_, err := b.bot.Request(setCommands)

	if err != nil {
		return err
	}

	return nil
}
