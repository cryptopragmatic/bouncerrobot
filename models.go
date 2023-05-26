package main

import (
	"errors"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	TelegramId int
	Votes      []Vote
}

func (m *Message) voteDel(userID uint) error {
	v := &Vote{}
	db.First(v, &Vote{MessageID: m.ID, UserID: userID})
	if v.ID == 0 {
		db.Model(m).Association("Votes").Append(&Vote{MessageID: m.ID, UserID: userID})
		return nil
	} else {
		return errors.New("Already voted.")
	}
}

func getMessage(mID int) *Message {
	m := &Message{}

	db.FirstOrCreate(m, &Message{TelegramId: mID})

	return m
}

type Vote struct {
	gorm.Model
	MessageID uint
	UserID    uint
}

type User struct {
	gorm.Model
	TelegramId int64
	Dels       int
}

func getUser(uID int64) *User {
	u := &User{}

	db.FirstOrCreate(u, &User{TelegramId: uID})

	return u
}
