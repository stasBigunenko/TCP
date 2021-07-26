package service

//server's actions

import (
	"log"
	"os"
	"strconv"
	"time"
)

func Help(msg string) string {
	msgNew := "You can choose the following commands: Hi, Bye, Time, Id"
	return msgNew
}

func Hi(msg string) string {
	msgNew := "Hello my friend!"
	return msgNew
}

func Bye(msg string) string {
	msgNew := "Good bye my friend"
	return msgNew
}

func Unknown() string {
	return "Unknown command, sorry. Please try again."
}

func Time(msg string) string {
	t := time.Now()
	msgNew := (t.Format("2006-01-02 15:04:05"))
	return msgNew
}

func IdConnection(i int) string {
	msgNew := "Your connection is number " + strconv.Itoa(i)
	return msgNew
}

//save func to the file with name clientCommand.txt
func SaveClientMsg(msg string, i int) {
	f, err := os.OpenFile("../../pkg/storage/clientCommand.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	t := time.Now()
	time := (t.Format("2006-01-02 15:04:05"))

	k := strconv.Itoa(i) //ID connection number

	if _, err := f.WriteString(msg + " ---------have been written in " + time + " by ID = " + k + "\n"); err != nil {
		log.Println(err)
	}
}

//save to the file with unknownCommand
func SaveNewCommand(msgNew string, i int) {
	f, err := os.OpenFile("../../pkg/storage/unknownCommand.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	t := time.Now()
	time := (t.Format("2006-01-02 15:04:05"))

	k := strconv.Itoa(i)

	if _, err := f.WriteString(msgNew + " ---------have been written in " + time + " by ID = " + k + "\n"); err != nil {
		log.Println(err)
	}
}
