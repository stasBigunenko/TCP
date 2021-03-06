package main

//client

import (
	"TCP/cmd/config"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	i := 0 // connection id
	// connect to server
	for {
		//connect:
		config := config.Set()

		conn, errConn := net.Dial(config.Protocol, config.Host+config.Port)
		if errConn != nil {
			continue
		} else {
			i++
		}
		if i <= 1 {
			log.Println("Connected to server...")
			fmt.Println("---")
		}
		for {
			// read message from stdin
			scannerStdin := bufio.NewScanner(os.Stdin)
			fmt.Print("Please write your command to the server: ")
			for scannerStdin.Scan() {
				text := scannerStdin.Text()
				fmt.Println("---")
				// send message to server
				_, errWrite := fmt.Fprintf(conn, text+"\n")
				if errWrite != nil {
					log.Println("Server have been terminated due to timeout.")
				}
				log.Print("Server receives: " + text)
				break
			}
			// listen for respond from server
			scannerConn := bufio.NewScanner(conn)
			for scannerConn.Scan() {
				log.Println("Server sends: " + scannerConn.Text())
				break
			}
			if errReadConn := scannerStdin.Err(); errReadConn != nil {
				log.Printf("Read error: %T %+v", errReadConn, errReadConn)
				return
			}
			fmt.Println("---") //separate messages
		}
	}
}
