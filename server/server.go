package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp4", ":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 9000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connecting to client:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Closing connection")
		conn.Close()
	}()

	for {
		// Reading from the client
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		message := strings.TrimRight(string(buf[:n]), "\n")
		fmt.Println("Received from client -->", message)

		if strings.TrimSpace(message) == "STOP" {
			fmt.Println("Disconnecting server...")
			_, err = conn.Write([]byte("Good Bye"))
			if err != nil {
				fmt.Println("Error")
				return
			}
            return
		}

		_, err = conn.Write([]byte(strings.TrimRight(message, "\n")))
		if err != nil {
			fmt.Println("Error sending message to client:", err)
			return
		}
		fmt.Println("Sent to client -->", message)
	}
}
