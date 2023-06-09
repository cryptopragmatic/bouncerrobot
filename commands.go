package main

import (
	"gopkg.in/telebot.v3"
)

func initCommands() {
	bot.Handle("/del", delCommand)
	bot.Handle("/start", startCommand)
}

func delCommand(c telebot.Context) error {
	if c.Message().IsReply() {
		err := voteDel(c.Message().ReplyTo.ID, c.Message().ReplyTo.Sender.ID, c.Sender().ID)
		if err != nil {
			bot.Send(c.Chat(), mAlreadyVoted)
			return err
		}

		del, err := checkDelete(c.Message().ReplyTo.ID, c.Message().ReplyTo.Sender.ID, c.Message().ReplyTo)
		if err != nil {
			bot.Send(c.Chat(), mNotAdmin)
		} else if del {
			bot.Send(c.Chat(), mDeleted)
		}

		ban := checkBan(c.Chat(), c.Message().ReplyTo.Sender.ID, c.Message().ReplyTo.Sender)
		if ban {
			bot.Send(c.Chat(), mBanned)
		}
	} else {
		bot.Send(c.Chat(), mMustReply)
	}

	return nil
}

func startCommand(c telebot.Context) error {
	bot.Send(c.Chat(), mWelcome)

	return nil
}
