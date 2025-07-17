package main

import (
	"filesword/middleware"
	"filesword/service"
	"filesword/utils"
	"log"
	"net"
	"strings"
)

func main() {
	config := utils.LoadConfig("./config.ini")
	server := service.NewTCP(config.Port)

	
	server.Use(middleware.Auth)

	server.Handle(func(conn net.Conn, data string) {
		addr := conn.RemoteAddr().String()

		if strings.HasPrefix(data, "@") {
			parts := strings.SplitN(data[1:], ":", 2)
			if len(parts) == 2 {
				target := strings.TrimSpace(parts[0])
				message := strings.TrimSpace(parts[1])

				server.SendTo(target, message)
				return
			}
		}

		server.Broadcast(data, addr)
	})

	if err := server.Start(); err != nil {
		panic(err)
	}

	log.Println("Veritabanı bağlantısı başarılı.")
}
