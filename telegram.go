package main

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func initTelegramBot() *telebot.Bot {
	b, err := telebot.NewBot(telebot.Settings{
		Token:     conf.TelegramAPIKey,
		Poller:    &telebot.LongPoller{Timeout: PollerTimeout * time.Second},
		Verbose:   false,
		ParseMode: "html",
	})

	if err != nil {
		log.Fatal(err)
	}

	return b
}

func logTelegram(message string) {
	message = "BouncerRobot:" + getCallerInfo() + message
	rec := &telebot.Chat{
		ID: int64(conf.LogAccount),
	}
	bot.Send(rec, message)
}

func notificationTelegram(message string) {
	rec := &telebot.Chat{
		ID: int64(conf.LogAccount),
	}
	bot.Send(rec, message)
}
