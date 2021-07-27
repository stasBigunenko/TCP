package tpcserver

import (
	"TCP/service"
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func HandleServerConnection(conn net.Conn, i int) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Closing client %v connection returned error: %v\n", i, err)
		}
	}()

	timeoutDuration := 10 * time.Second

	// receive the message
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		// Set a deadline for reading. Read operation will fail if no data
		// is received after deadline.
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		msg := scanner.Text()
		service.SaveClientMsg(msg, i)
		msg = strings.ToLower(msg)
		var msgNew string
		switch {
		case msg == "help":
			msgNew = service.Help()
		case msg == "hi":
			msgNew = service.Hi()
		case msg == "bye":
			msgNew = service.Bye()
		case msg == "time":
			msgNew = service.Time()
		case msg == "id":
			msgNew = service.IdConnection(i)
		default:
			msgNew = service.Unknown()
			service.SaveNewCommand(msg, i)
		}

		//respond the answer from the server to the client
		if _, err := conn.Write([]byte(msgNew + "\n")); err != nil {
			log.Printf("Client %v has closed writer: %v\n", i, err)
			break
		}
		log.Printf("Client %v\n received: %v\n returned: %v\n", i, msg, msgNew)
	}
	log.Printf("Client %v disconnected...\n", i)
}
