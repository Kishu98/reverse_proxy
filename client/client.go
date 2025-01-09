package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:9002")
	if err != nil {
		fmt.Println("Error connecting to the reverse proxy:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to the reverse proxy at:", conn.RemoteAddr())

	user_input_reader := bufio.NewReader(os.Stdin)
	for {

		// Read the input from the stdin
		fmt.Print("--> ")
		user_input, err := user_input_reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading client")
			break
		}

		if strings.TrimSpace(user_input) == "STOP" {
			fmt.Println("Disconnecting...")
            _, err = conn.Write([]byte(user_input))
            if err != nil {
                fmt.Println("Error")
            }
			break
		}

		// Send the input from the user to the reverse proxy
		_, err = conn.Write([]byte(user_input))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			break
		}

		// Receive the message sent from the server through the reverse proxy
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from the server:", err)
				break
			}
		}

		// Print the message received
		receivedMsg := string(buf[:n])
		fmt.Println("Received -->", receivedMsg)
	}
}
