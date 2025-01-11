package Server

import (
	"net"
)

type Server struct {
	Listener net.Listener
}

func NewServer() (Server, error) {
	listener, err := net.Listen("tcp4", "")
	if err != nil {
		return Server{}, err
	}

	return Server{Listener: listener}, nil
}
