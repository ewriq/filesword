package middleware

import (
	"filesword/model"

	"fmt"
	"net"
)

func Logger(next model.HandlerFunc)  model.HandlerFunc {
	return func(conn net.Conn, data string) {
		fmt.Printf("📥 Gelen veri: %s\n", data)
		fmt.Printf("📥 Bağlantı adresi: %s\n", conn.RemoteAddr().String())
		fmt.Printf("📥 Bağlantı tipi: %T\n", conn)
		fmt.Println(data)
		fmt.Println("----------------------------------")
		
		next(conn, data)
	}
}
