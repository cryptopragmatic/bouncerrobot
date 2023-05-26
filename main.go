package main

import (
	"log"

	"gopkg.in/telebot.v3"
)

var conf *Config

var bot *telebot.Bot

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	bot = initTelegramBot()

	log.Println("VoteDelBot started. 🚀")

	bot.Start()
}
