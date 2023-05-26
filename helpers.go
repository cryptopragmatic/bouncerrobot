package main

import (
	"fmt"
	"path"
	"runtime"

	"gopkg.in/telebot.v3"
)

func getCallerInfo() (info string) {
	// pc, file, lineNo, ok := runtime.Caller(2)
	_, file, lineNo, ok := runtime.Caller(2)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	// funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // The Base function returns the last element of the path
	return fmt.Sprintf("%s:%d: ", fileName, lineNo)
}

func voteDel(messageID int, ownerID int64, voterID int64) error {
	m := getMessage(messageID)
	getUser(ownerID)
	return m.voteDel(voterID)
}

func checkDelete(messageID int, msg *telebot.Message) (bool, error) {
	m := getMessage(messageID)
	count := db.Model(m).Association("Votes").Count()
	if count == conf.Votes {
		err := bot.Delete(msg)
		return true, err
	}
	return false, nil
}

func checkBan(userID int64) {

}
