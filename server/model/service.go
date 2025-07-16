package model

import "net"

type HandlerFunc func(conn net.Conn, data string)
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
