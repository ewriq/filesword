package middleware

import (
	"net"
	"strings"
	"sync"

	"filesword/model"
	"filesword/utils"
)

var (
	authenticatedConns = make(map[net.Conn]bool)
	mu                 sync.Mutex
)

func Auth(next model.HandlerFunc) model.HandlerFunc {
	config := utils.LoadConfig("./config.ini")
	username := config.Username;
	password := config.Password;
	return func(conn net.Conn, data string) {
		mu.Lock()
		auth := authenticatedConns[conn]
		mu.Unlock()

		if !auth {
			parts := strings.SplitN(data, ":", 2)
			if len(parts) != 2 {
				conn.Write([]byte("kullanıcı adı ve şifre formatı: admin:1234\n"))
				conn.Close()
				return
			}

			user := strings.TrimSpace(parts[0])
			pass := strings.TrimSpace(parts[1])

			if user != username || pass != password {
				conn.Write([]byte("geçersiz giriş\n"))
				conn.Close()
				return
			}

			mu.Lock()
			authenticatedConns[conn] = true
			mu.Unlock()

			conn.Write([]byte("✅ Giriş başarılı. Mesaj gönderebilirsin.\n"))
			return
		}

		next(conn, data)
	}
}
