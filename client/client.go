package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Provide the hostname and port")
		os.Exit(1)
	}

	connect_to := args[1]
	conn, err := net.Dial("tcp", connect_to)
	if err != nil {
		log.Println("Error connecting to server")
		os.Exit(1)
	}

	fmt.Println("Connected to server")

    // Here we are accepting the message from the stdin, which is message
    // There are two readers in the client. Wow
	input_message := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>")
        message, err := input_message.ReadString('\n')
		if err != nil {
			log.Println("Error receiving message:", err)
			os.Exit(1)
		}
		fmt.Fprintf(conn, "%s\n", message)

        // This is the second reader accepting message from the server
        input_message_server, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            log.Println("Error receiving message from the server:", err)
            os.Exit(1)
        }
        fmt.Print("->", input_message_server)

        if strings.TrimSpace(message) == "STOP" {
            fmt.Println("Exiting...")
            return
        }
	}
}
