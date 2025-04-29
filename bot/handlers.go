package tgBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	START   = "start"
	HELP    = "help"
	TRACK   = "track"
	UNTRACK = "untrack"
	LIST    = "list"
)

var description = map[string]string{
	START:   "start",
	HELP:    "help",
	TRACK:   "track",
	UNTRACK: "untrack",
	LIST:    "list",
}

var commands = []tgbotapi.BotCommand{
	{
		Command:     START,
		Description: description[START],
	},
	{
		Command:     HELP,
		Description: description[HELP],
	},
	{
		Command:     TRACK,
		Description: description[TRACK],
	},
	{
		Command:     UNTRACK,
		Description: description[UNTRACK],
	},
	{
		Command:     LIST,
		Description: description[LIST],
	},
}

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Init text msg")

	if err := b.handlerShowDownList(); err != nil {
		return err
	}

	switch message.Command() {
	case START:
		return b.handlerStart(msg)
	case HELP:
		return b.handlerHelp(msg)
	case TRACK:
		return b.handlerTrack(msg)
	case UNTRACK:
		return b.handlerUntrack(msg)
	case LIST:
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

func (b *Bot) handlerStart(msg tgbotapi.MessageConfig) error {
	msg.Text = fmt.Sprintf("%s and %s", msg.ChatID)
	//userId := msg.ChannelUsername
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handlerHelp(msg tgbotapi.MessageConfig) error {
	list := ""
	for i := 0; i < len(commands); i++ {
		list += "/" + commands[i].Command + " - " + commands[i].Description + "\n"
	}
	msg.Text = list
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
	fmt.Println(commands[0].Command)
	_, err := b.bot.Send(msg)
	return err
}
