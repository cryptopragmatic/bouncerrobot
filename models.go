package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TelegramId int
	Votes      int
}

func (m *Message) voteDel() {
	m.Votes++
	db.Save(m)
}

func getMessage(mID int) *Message {
	m := &Message{}

	db.FirstOrCreate(m, &Message{TelegramId: mID})

	return m
}

type User struct {
	gorm.Model
	TelegramId int64
	Dels       int
}
