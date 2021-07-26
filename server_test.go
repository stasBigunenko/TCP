package server_test

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestConn(t *testing.T) {
	msg := "Hi!"

	go func() {
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, msg); err != nil {
			t.Fatal(err)
		}
	}()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		scanner.Scan()

		fmt.Println(msg)
		if msgNew := scanner.Text(); msgNew != msg {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msgNew, msg)
		}
		return // Done
	}
}
