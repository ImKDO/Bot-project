package tgBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	Start   = "start"
	Help    = "help"
	Track   = "track"
	Untrack = "untrack"
	List    = "list"
)

var commands = []tgbotapi.BotCommand{
	{
		Command:     Start,
		Description: "Start the bot",
	},
	{
		Command:     Help,
		Description: "Show help",
	},
	{
		Command:     Track,
		Description: "Track the bot",
	},
	{
		Command:     Untrack,
		Description: "Untrack the bot",
	},
	{
		Command:     List,
		Description: "List all bots",
	},
}

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Init text msg")

	if err := b.handlerShowDownList(); err != nil {
		return err
	}

	switch message.Command() {
	case Start:
		return b.handlerStart(msg)
	case Help:
		return b.handlerHelp(msg)
	case Track:
		return b.handlerTrack(msg)
	case Untrack:
		return b.handlerUntrack(msg)
	case List:
		return b.handlerList(msg)
	default:
		return nil
	}
}
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

func (b *Bot) handlerShowDownList() error {
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	_, err := b.bot.Request(setCommands)

	if err != nil {
		return err
	}

	return nil

}

func (b *Bot) handlerDescriptCommand(command string) string {
	return command
}
func (b *Bot) handlerStart(msg tgbotapi.MessageConfig) error {
	msg.Text = "Команда /start"
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handlerHelp(msg tgbotapi.MessageConfig) error {
	msg.Text = "Команда /help"
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handlerTrack(msg tgbotapi.MessageConfig) error {
	msg.Text = "Команда /track"
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handlerUntrack(msg tgbotapi.MessageConfig) error {
	msg.Text = "Команда /untrack"
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handlerList(msg tgbotapi.MessageConfig) error {
	msg.Text = "Команда /list"
	_, err := b.bot.Send(msg)
	return err
}
