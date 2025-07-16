package service

import (
	"fmt"
	"net"
)

func Handler(conn net.Conn, data string) {
	response := fmt.Sprintf("✅ Sunucu aldı: %s\n", data)
	conn.Write([]byte(response))
}
