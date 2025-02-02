package internal

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/ARF-DEV/ping-pong-mp/common/core"
	"github.com/ARF-DEV/ping-pong-mp/common/network"
)

type Server struct {
	scene *core.Scene
	conn  net.Conn
	q     chan network.InputMessage
	t     int32
}

func NewServer(conn net.Conn) *Server {
	s := Server{}
	s.conn = conn
	s.q = make(chan network.InputMessage, 30)
	s.scene = core.CreateGame("server")

	return &s
}

func (s *Server) ProcessConn() {
	// b := bufio.NewReader(s.conn)
	timer := time.NewTimer(1 * time.Minute)
	i := 0
ExitLoop:
	for {
		select {
		case <-timer.C:
			fmt.Println("test a")
			break ExitLoop
		default:
			i++
			b := make([]byte, 1024)
			n, err := s.conn.Read(b)
			if err != nil {
				continue
			}
			ins := s.getInputMsg(b[:n])
			fmt.Println(string(b[:n]))
			for _, in := range ins {
				s.q <- in
			}
			timer.Reset(1 * time.Minute)
		}
	}
	fmt.Println("waduh")
}

func (s *Server) ProcessInput() {
	tick := time.NewTicker((1 * time.Second) / network.SERVER_TICK)
	timer := time.NewTimer(1 * time.Minute)
OuterLoop:
	for {
		select {
		case <-tick.C:
			s.scene.UpdateFromNonInput()
		InnerLoop:
			for {
				select {
				case in := <-s.q:
					s.scene.UpdateFromInput(in.Input)
					fmt.Println("client tick: ", in.Tick)
					// TBD: should s.t++ be here also??
					s.t++
					break InnerLoop
				default:
					break InnerLoop
				}
			}
			state := s.scene.GetServerSceneState()
			state.Tick = s.t
			fmt.Println("server tick: ", s.t)
			// utils.PrintToJSON(state.Actors)
			network.Send(s.conn, state)
			timer.Reset(1 * time.Minute)
			s.t++
			// if s.t == 50 {
			// 	panic("adkaowdk")
			// }
			// panic("awodawodk")

		case <-timer.C:
			break OuterLoop
		}
	}
}

func (s *Server) getInputMsg(buf []byte) []network.InputMessage {
	res := []network.InputMessage{}
	splitStr := strings.Split(string(buf), "\n")
	for _, str := range splitStr {
		msg := network.InputMessage{}
		json.Unmarshal([]byte(str), &msg)
		res = append(res, msg)

	}
	return res
}
