package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"time"

	"github.com/ARF-DEV/ping-pong-mp/common/core"
	"github.com/ARF-DEV/ping-pong-mp/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// rl.InitWindow(800, 450, "raylib [core] example - basic window")
	// defer rl.CloseWindow()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("server listening at port 8080")
	conn, err := listener.Accept()
	if err != nil {
		log.Printf("error when accepting connection: %v", err)
		return
	}
	game := core.CreateGame("server")
	rl.SetTargetFPS(60)
	for {
		b := bufio.NewReader(conn)
		for {
			d := core.InputMessage{}
			str, err := b.ReadString('\n')
			time.Sleep(100 * time.Millisecond) // simulate travel time between server and client
			json.Unmarshal([]byte(str), &d)
			utils.PrintToJSON(d)
			game.UpdateFromInput(d.Input)
			response, _ := json.Marshal(game.GetSceneState())
			conn.Write(response)
			if err == io.EOF {
				break
			}
		}
	}

}
