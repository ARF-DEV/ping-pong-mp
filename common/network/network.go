package network

import (
	"encoding/json"
	"net"
)

func Send(dst net.Conn, data any) {
	payload, _ := json.Marshal(data)
	dst.Write(payload)
}

type InputMessage struct {
	Input int32
}
