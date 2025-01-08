package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Enter the port number:")
		os.Exit(1)
	}

	port := ":" + args[1]
	conn, err := net.Listen("tcp4", port)
	if err != nil {
		log.Printf("Error starting server on %v: %v", conn.Addr(), err)
		return
	}
	defer conn.Close()

	for {
		serverConn, err := conn.Accept()
		if err != nil {
			log.Println("Error connecting to client:", err)
			return
		}
		log.Println("Connected to client")

		go handleConnection(serverConn)
	}
}

func handleConnection(c net.Conn) {
	defer func() {
		c.Close()
	}()

	for {
		client_message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println("Error reading message from client:", err)
			return
		}
		fmt.Println(">", client_message)

		fmt.Fprintf(c, "From server: %s", client_message)
	}
}
