package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Server struct {
	listenAddr string
	listener   net.Listener
}

func createServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
	}
}

func (s *Server) connect() {

	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		log.Fatal(err)
	}

	go s.accept()
	s.listener = ln
}

func (s *Server) read(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	for {

		n, err := conn.Read(buff)

		if err != nil {
			fmt.Printf("Read error : %v", err)
			os.Exit(1)
		}

		msg := buff[:n]
		fmt.Printf("message is : %v \n", strings.TrimSpace(string(msg)))
	}
}

func (s *Server) accept() {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			fmt.Printf("error : %v", err)
			os.Exit(1)
		}
		go s.read(conn)
		fmt.Printf("Connected remote user : %v", conn.RemoteAddr().String())
	}

}
func main() {

	server := createServer(":8080")
	server.connect()

	defer server.listener.Close()

	select {}

}
