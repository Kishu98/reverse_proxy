package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp4", ":9002")
	if err != nil {
		fmt.Println("Error starting the reverse proxy server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Reverse proxy is listening at port 9002...")

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connecting to client:", err)
			continue
		}
		fmt.Println("Connected to client:", clientConn.RemoteAddr())

		go handleClientConnection(clientConn)
	}
}

func handleClientConnection(clientConn net.Conn) {
	serverConn, err := net.Dial("tcp4", "localhost:9000")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	fmt.Println("Connected to the server at:", serverConn.RemoteAddr())
    defer func() {
        fmt.Println("Closing connections")
        clientConn.Close()
        serverConn.Close()
    }()

	go forwardData(clientConn, serverConn)
	forwardData(serverConn, clientConn)

}

func forwardData(src, dest net.Conn) {
	_, err := io.Copy(dest, src)
	if err != nil {
		fmt.Println("Error writing to server from client:", err)
	}
}
