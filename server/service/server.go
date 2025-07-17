package service

import (
	"bufio"
	"fmt"
	"net"

	"sync"

	"filesword/model"
)

type TCPServer struct {
	address     string
	middlewares []model.MiddlewareFunc
	handler     model.HandlerFunc
	clients     map[string]net.Conn
	mu          sync.Mutex
}

func NewTCP(addr string) *TCPServer {
	return &TCPServer{
		address: addr,
		clients: make(map[string]net.Conn),
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

	fmt.Println("🚀 Sunucu başlatıldı", s.address, "portunda")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("⚠️ Bağlantı hatası:", err)
			continue
		}

		addr := conn.RemoteAddr().String()

		s.mu.Lock()
		s.clients[addr] = conn
		s.mu.Unlock()

		fmt.Println("🟢 Bağlanan:", addr)

		go s.handleClient(conn, addr)
	}
}

func (s *TCPServer) handleClient(conn net.Conn, addr string) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, addr)
		s.mu.Unlock()
		fmt.Println("🔴 Ayrıldı:", addr)
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		s.handler(conn, text) 
	}
}

func (s *TCPServer) SendTo(target string, message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if conn, ok := s.clients[target]; ok {
		conn.Write([]byte(message + "\n"))
	}
}

func (s *TCPServer) Broadcast(message, sender string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for addr, conn := range s.clients {
		if addr != sender {
			conn.Write([]byte(message + "\n"))
		}
	}
}
