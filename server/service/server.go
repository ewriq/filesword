package service

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"filesword/model"
)


type TCPServer struct {
	address     string
	middlewares []model.MiddlewareFunc
	handler     model.HandlerFunc
}

func NewTCP(addr string) *TCPServer {
	return &TCPServer{
		address: addr,
	}
}

func (s *TCPServer) Use(mw model.MiddlewareFunc) {
	s.middlewares = append(s.middlewares, mw)
}

func (s *TCPServer) Handle(h model.HandlerFunc) {
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		h = s.middlewares[i](h)
	}
	s.handler = h
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("🔌 TCP server listening on", s.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Bağlantı hatası:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("🔌 Bağlantı kapandı")
			return
		}
		data = strings.TrimSpace(data)
		s.handler(conn, data)
	}
}
