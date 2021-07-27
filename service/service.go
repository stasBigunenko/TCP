package service

//server's actions

import (
	"log"
	"os"
	"strconv"
	"time"
)

func Help() string {
	return "You can choose the following commands: Hi, Bye, Time, Id"
}

func Hi() string {
	return "Hello my friend!"
}

func Bye() string {
	return "Good bye my friend"
}

func Unknown() string {
	return "Unknown command, sorry. Please try again."
}

func Time() string {
	t := time.Now()
	msgNew := (t.Format("2006-01-02 15:04:05"))
	return msgNew
}

func IdConnection(i int) string {
	return "Your connection is number " + strconv.Itoa(i)
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
func SaveNewCommand(msg string, i int) {
	f, err := os.OpenFile("../../pkg/storage/unknownCommand.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	t := time.Now()
	time := (t.Format("2006-01-02 15:04:05"))

	k := strconv.Itoa(i)

	if _, err := f.WriteString(msg + " ---------have been written in " + time + " by ID = " + k + "\n"); err != nil {
		log.Println(err)
	}
}
