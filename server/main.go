package main

import (
	"filesword/middleware"
	"filesword/service"
	"log"
)

func main() {
	server := service.NewTCP(":9000")

	server.Use(middleware.LoggerMiddleware)
	server.Use(middleware.UppercaseMiddleware)


	server.Handle(service.Handler)

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
