package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TelegramId int64
	Votes      int
}

type User struct {
	gorm.Model
	TelegramId int64
	Dels       int
}
