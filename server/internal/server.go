package internal

import (
	"bufio"
	"encoding/json"
	"net"
	"time"

	"github.com/ARF-DEV/ping-pong-mp/common/core"
	"github.com/ARF-DEV/ping-pong-mp/common/network"
)

const (
	SERVER_TICK = 33
)

type Server struct {
	scene *core.Scene
	conn  net.Conn
	q     chan network.InputMessage
}

func NewServer(conn net.Conn) *Server {
	s := Server{}
	s.conn = conn
	s.q = make(chan network.InputMessage, 30)
	s.scene = core.CreateGame("server")

	return &s
}

func (s *Server) ProcessConn() {
	b := bufio.NewReader(s.conn)
	timer := time.NewTimer(1 * time.Minute)

ExitLoop:
	for {
		select {
		case <-timer.C:
			break ExitLoop
		default:
			for {
				d := network.InputMessage{}
				str, err := b.ReadString('\n')
				if err != nil {
					break
				}
				time.Sleep(100 * time.Millisecond) // simulate travel time between server and client
				json.Unmarshal([]byte(str), &d)
				s.q <- d
				timer.Reset(1 * time.Minute)
			}
		}
	}
}

func (s *Server) ProcessInput() {
	tick := time.NewTicker(1 / SERVER_TICK)
	timer := time.NewTimer(1 * time.Minute)
OuterLoop:
	for {
		select {
		case <-tick.C:
		InnerLoop:
			for {
				select {
				case in := <-s.q:
					s.scene.UpdateFromInput(in.Input)
				default:
					break InnerLoop
				}
			}
			state := s.scene.GetSceneState()
			network.Send(s.conn, state)
			timer.Reset(1 * time.Minute)

		case <-timer.C:
			break OuterLoop
		}
	}
}
