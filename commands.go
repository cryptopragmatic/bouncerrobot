package main

import (
	"gopkg.in/telebot.v3"
)

func initCommands() {
	bot.Handle("/del", delCommand)
}

func delCommand(c telebot.Context) error {
	if c.Message().IsReply() {
		voteDel(c.Message().ReplyTo.ID)

		checkDelete(c.Message().ReplyTo.ID, c.Message().ReplyTo.Sender.ID)

		checkBan(c.Message().ReplyTo.Sender.ID)
	} else {
		bot.Send(c.Chat(), mMustReply)
	}

	return nil
}
