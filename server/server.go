package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Server listener at port 9000
	listener, err := net.Listen("tcp4", ":9000")
	if err != nil {
		log.Println("Error starting server")
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening at port 9000")

	client := 0
	for {
		// Accepting connection to the server
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting client:", conn.RemoteAddr())
			continue
		}
		fmt.Printf("Connected to client %d: %v", client, conn.RemoteAddr())
		client++

		go handleConnection(conn, client-1)
	}
}

func handleConnection(conn net.Conn, client int) {
	defer func() {
		fmt.Printf("Disconnecting from client %d\n", client)
		defer conn.Close()
	}()
	for {
		// Receiving message by the read command and storing the values in buf
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading from client")
			break
		}
		message := strings.TrimSpace(string(buf[:n]))
		fmt.Printf("Client %d: %s\n", client, message)

		// Disconnecting if user sends STOP
		if message == "STOP" {
			break
		}

		//Sending same message back
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to client %d: %v", client, err)
			break
		}
		fmt.Printf("To client %d: %s\n", client, message)

	}
}
