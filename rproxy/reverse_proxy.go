package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	proxyListener, err := net.Listen("tcp4", "")
	if err != nil {
		log.Println("Error starting proxy:", err)
		os.Exit(1)
	}
	log.Println("Proxy listening at:", proxyListener.Addr())
	defer proxyListener.Close()

	for {
		clientConn, err := proxyListener.Accept()
		if err != nil {
			log.Println("Error connecting to client:", err)
			break
		}
		log.Println("Proxy connected to client at", clientConn.RemoteAddr())

		go handleClientConnection(clientConn)
	}
}

func handleClientConnection(clientConn net.Conn) {
	serverListener, err := net.Listen("tcp4", "")
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}
	log.Println("Server listening at:", serverListener.Addr())
	// defer serverListener.Close()

	go func() {
		serverConn, err := serverListener.Accept()
		if err != nil {
			log.Println("Error connecting to the proxy:", err)
			return
		}
		log.Println("Server connected to proxy at", serverConn.RemoteAddr())

		defer serverConn.Close()
		for {
			buf := make([]byte, 2048)
			n, err := serverConn.Read(buf)
			if err != nil {
				log.Println("Error reading data from the proxy:", err)
				break
			}
			log.Println("Received message", string(buf[:n]))
		}

	}()

	serverConn, err := net.Dial("tcp4", serverListener.Addr().String())
	if err != nil {
		log.Println("Error connecting to the server:", err)
		return
	}
	log.Println("Proxy connected to server at", serverConn.RemoteAddr())

	go forwardData(clientConn, serverConn)
	forwardData(serverConn, clientConn)
}

func forwardData(src, dest net.Conn) {
	log.Println("Sending to:", dest.RemoteAddr())
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Println("Error writing to:", dest.RemoteAddr())
		return
	}
	log.Println("Sent Message from :", src.RemoteAddr())
	log.Println("Received message from:", dest.RemoteAddr())
}
