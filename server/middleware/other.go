package middleware

import (
	"filesword/database"
	"filesword/model"

	"fmt"
	"net"
)

var err error

func Logger(next model.HandlerFunc)  model.HandlerFunc {
	return func(conn net.Conn, data string) {
		fmt.Printf("📥 Gelen veri: %s\n", data)
		fmt.Printf("📥 Bağlantı adresi: %s\n", conn.RemoteAddr().String())
		fmt.Println("----------------------------------")
		err = database.Log(data, "message")
		if err != nil {
			fmt.Printf("❌ Loglama hatası: %v\n", err)
		}
		err = database.Log(conn.RemoteAddr().String(), "address")
		if err != nil {
			fmt.Printf("❌ Loglama hatası: %v\n", err)
		}
		next(conn, data)
	}
}
