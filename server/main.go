package main

import (
	"filesword/middleware"
	"filesword/service"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	server := service.NewTCP(":9000")

	server.Use(middleware.LoggerMiddleware)

	server.Handle(func(conn net.Conn, data string) {
		addr := conn.RemoteAddr().String()

		if strings.HasPrefix(data, "@") {
			parts := strings.SplitN(data[1:], ":", 2)
			if len(parts) == 2 {
				target := strings.TrimSpace(parts[0])
				message := strings.TrimSpace(parts[1])
				server.SendTo(target, fmt.Sprintf("%s", message))
				return
			}
		}

		server.Broadcast(fmt.Sprintf("%s", data), addr)
	})

	if err := server.Start(); err != nil {
		panic(err)
	}
	
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
