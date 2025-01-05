package main

import (
	"log"
	"net"

	"github.com/ARF-DEV/ping-pong-mp/server/internal"
)

func main() {
	// rl.InitWindow(800, 450, "raylib [core] example - basic window")
	// defer rl.CloseWindow()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("server listening at port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error when accepting connection: %v", err)
			return
		}
		s := internal.NewServer(conn)
		go s.ProcessConn()
		go s.ProcessInput()
	}

}
