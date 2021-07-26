package tpcserver

import (
	"../../service"
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

	timeoutDuration := 20 * time.Second

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
		case string(msg) == "help":
			msgNew = service.Help(msg)
		case string(msg) == "hi":
			msgNew = service.Hi(msg)
		case string(msg) == "bye":
			msgNew = service.Bye(msg)
		case string(msg) == "time":
			msgNew = service.Time(msg)
		case string(msg) == "id":
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
