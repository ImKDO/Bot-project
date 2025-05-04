package tgBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handlerStart(msg *tgbotapi.Message) error {
	greatMsg := msg.From.FirstName + " " + msg.From.LastName
	msgText := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("%s", greatMsg))

	_, err := b.bot.Send(msgText)
	return err
}

var stateUsers = make(map[int64]string)
var timeToken = make(map[int64]string)

func (b *Bot) handlerHelp(msg *tgbotapi.Message) error {
	list := ""
	commands := b.setCommands(description)
	for i := 0; i < len(commands); i++ {
		list += "/" + commands[i].Command + " - " + commands[i].Description + "\n"
	}
	msgText := tgbotapi.NewMessage(msg.Chat.ID, list)
	_, err := b.bot.Send(msgText)
	return err
}

func (b *Bot) handlerTrack(msg *tgbotapi.Message) error {
	stateUsers[msg.Chat.ID] = "awaiting link"
	timeToken[msg.Chat.ID] = msg.Text

	msgText := tgbotapi.NewMessage(msg.Chat.ID, "Введите ссылку для отслеживания")
	_, err := b.bot.Send(msgText)
	if err != nil {
		return err
	}
	msgText = tgbotapi.NewMessage(msg.Chat.ID, msg.Text)

	//_, err = b.bot.Send()
	return err
}

//func (b *Bot) handlerUntrack(msg *tgbotapi.Message) error {
//	msg.Text = "Команда /untrack"
//	_, err := b.bot.Send(msg)
//	return err
//}
//func (b *Bot) handlerList(msg *tgbotapi.Message) error {
//	fmt.Println(commands[0].Command)
//	_, err := b.bot.Send(msg)
//	return err
//}

func (b *Bot) handleCommands(message *tgbotapi.Message) error {

	if err := b.handlerShowDownList(); err != nil {
		return err
	}

	switch message.Command() {
	case START:
		return b.handlerStart(message)
	case HELP:
		return b.handlerHelp(message)
	case TRACK:
		return b.handlerTrack(message)
	//case UNTRACK:
	//	return b.handlerUntrack(message)
	//case LIST:
	//	return b.handlerList(message)
	default:
		return nil
	}
}
