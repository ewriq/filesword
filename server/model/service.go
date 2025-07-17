package model

import "net"

type HandlerFunc func(conn net.Conn, data string)
type MiddlewareFunc func(next HandlerFunc) HandlerFunc


type Meta struct {
	FileName string `json:"fileName"`
	Mode     uint32 `json:"mode"`
	AtimeMs  int64  `json:"atimeMs"`
	MtimeMs  int64  `json:"mtimeMs"`
	FileData string `json:"fileData"`
	Username string `json:"username"`
	Password string `json:"password"`
}
