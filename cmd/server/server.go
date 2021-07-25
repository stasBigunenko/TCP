package main

//server

import (
	"../../pkg/tpcserver"
	"../config"
	"fmt"
	"log"
	"net"
)

func main() {
	config := config.New()

	log.Println("Server launched...")
	// listen according the config data
	ln, err := net.Listen(config.Protocol, config.Port)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		i++
		log.Printf("Client %v connected...", i)
		fmt.Println("---")
		// handle the connection
		go tpcserver.HandleServerConnection(conn, i)
	}
}