package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	// Connecting to the Reverse Proxy
	conn, err := net.Dial("tcp4", args[1])
	if err != nil {
		fmt.Println("Error connecting to the reverse proxy:", err)
		os.Exit(1)
	}
	// Will close connection after the completion of the for loop
	defer conn.Close()
	fmt.Println("Connected to the reverse proxy at:", conn.RemoteAddr())

	// Reader for getting input from the stdin
	user_input_reader := bufio.NewReader(os.Stdin)

	for {
		// Read the input from the stdin
		fmt.Print("--> ")
		user_input, err := user_input_reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading client")
			break
		}

		// STOP will trigger closing of connection
		if strings.TrimSpace(user_input) == "STOP" {
			fmt.Println("Disconnecting...")
			_, err = conn.Write([]byte(user_input))
			if err != nil {
				fmt.Println("Error")
			}
			time.Sleep(time.Second)
			break
		}

		// Send the input from the user to the reverse proxy
		_, err = conn.Write([]byte(user_input))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			break
		}

		// Receive the message sent from the server through the reverse proxy
		// buf := make([]byte, 1024)
		// n, err := conn.Read(buf)
		// if err != nil {
		// 	if err != io.EOF {
		// 		fmt.Println("Error reading from the server:", err)
		// 		break
		// 	}
		// }
		//
		// // Print the message received
		// receivedMsg := string(buf[:n])
		// fmt.Println("Received -->", receivedMsg)
	}
}
