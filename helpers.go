package main

import (
	"fmt"
	"log"
	"path"
	"runtime"
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

func voteDel(messageID int, userID int64) error {
	m := getMessage(messageID)
	u := getUser(userID)
	return m.voteDel(u.ID)
}

func checkDelete(messageID int, userID int64) {
	m := getMessage(messageID)
	log.Println(m)
}

func checkBan(userID int64) {

}
