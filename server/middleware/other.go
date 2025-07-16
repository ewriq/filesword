package middleware

import (
	"filesword/model"

	"fmt"
	"net"
	"strings"
)

func LoggerMiddleware(next model.HandlerFunc)  model.HandlerFunc {
	return func(conn net.Conn, data string) {
		fmt.Printf("📥 Gelen veri: %s\n", data)
		next(conn, data)
	}
}

func UppercaseMiddleware(next  model.HandlerFunc)  model.HandlerFunc {
	return func(conn net.Conn, data string) {
		data = strings.ToUpper(data)
		next(conn, data)
	}
}
