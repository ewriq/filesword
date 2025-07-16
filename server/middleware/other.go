package middleware

import (
	"filesword/model"

	"fmt"
	"net"
)

func LoggerMiddleware(next model.HandlerFunc)  model.HandlerFunc {
	return func(conn net.Conn, data string) {
		fmt.Printf("📥 Gelen veri: %s\n", data)
		next(conn, data)
	}
}
