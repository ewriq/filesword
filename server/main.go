package main

import (
	"filesword/middleware"
	"filesword/service"
	"filesword/utils"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	config := utils.LoadConfig("./config.ini")
	
	server := service.NewTCP(config.Port)

	server.Use(middleware.Auth)
	server.Use(middleware.Logger)

	server.Handle(func(conn net.Conn, data string) {
		addr := conn.RemoteAddr().String()
		fmt.Printf("📥 Gelen veri: %s\n", data)
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
