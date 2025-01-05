package network

import (
	"encoding/json"
	"net"
)

func Send(dst net.Conn, data any) {
	payload, _ := json.Marshal(data)
	dst.Write(append(payload, '\n'))
}

type InputMessage struct {
	Input int32
}

const (
	SERVER_TICK = 60
)
