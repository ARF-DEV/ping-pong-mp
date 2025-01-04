package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	AreaWidth  float32 = 800
	AreaHeight float32 = 450
	PadSpeed   float32 = 150
	BallSpeed  float32 = 130
)

type Scene struct {
	Area rl.Rectangle
	// p1   *Player
	// p2   *Player
	// ball *Ball
	Actors []Actor
	conn   net.Conn
}

func CreateGame(t string) *Scene {
	mWidth := rl.GetScreenWidth()
	if mWidth == 0 {
		mWidth = int(AreaWidth)
	}
	fmt.Println(mWidth)
	mHeight := rl.GetScreenHeight()
	if mHeight == 0 {
		mHeight = int(AreaHeight)
	}
	fmt.Println(mHeight)
	mCenterX := mWidth / 2
	mCenterY := mHeight / 2
	areaTopX := float32(mCenterX) - AreaWidth/2
	areaTopY := float32(mCenterY) - AreaHeight/2
	pOneCenterY := AreaHeight/2 + areaTopY
	pOneCenterX := areaTopX + 20
	p1 := Player{
		Rect:    rl.Rectangle{X: pOneCenterX - 10, Y: float32(pOneCenterY) - 40, Width: 20.0, Height: 80.0},
		UpKey:   rl.KeyW,
		DownKey: rl.KeyS,
		Color:   rl.Black,
	}

	pTwoCenterY := AreaHeight/2 + areaTopY
	pTwoCenterX := areaTopX + AreaWidth - 20
	p2 := Player{
		Rect:    rl.Rectangle{X: pTwoCenterX - 10, Y: pTwoCenterY - 40, Width: 20, Height: 80},
		UpKey:   rl.KeyUp,
		DownKey: rl.KeyDown,
		Color:   rl.Blue,
	}

	ballPosX := AreaWidth/2 + areaTopX
	ballPosY := AreaHeight/2 + areaTopY
	ball := Ball{
		Pos:   rl.Vector2{X: float32(ballPosX), Y: float32(ballPosY)},
		Rad:   6,
		Color: rl.Red,
		Dir:   rl.Vector2{X: 0.5, Y: 0.1},
	}

	s := &Scene{
		Area: rl.NewRectangle(areaTopX, areaTopY, AreaWidth, AreaHeight),
	}
	s.AddActor(&p1, &p2, &ball)

	if t == "client" {
		var err error
		s.conn, err = net.Dial("tcp", ":8080")
		if err != nil {
			panic(err)
		}
		go ProcessConn(s.conn, s)
	}

	return s
}

func (s *Scene) AddActor(lst ...Actor) {
	s.Actors = append(s.Actors, lst...)
}
func (g *Scene) Update() {
	for i := range g.Actors {
		g.Actors[i].Update(g)
	}
}

func (g *Scene) Draw() {
	rl.DrawRectangleLinesEx(g.Area, 2, rl.Black)
	for i := range g.Actors {
		g.Actors[i].Draw()
	}
}
func (g *Scene) UpdateFromInput(in int32) {
	for i := range g.Actors {
		g.Actors[i].UpdateFromInput(in)
	}
	// utils.PrintToJSON(g.Actors)
}

type InputMessage struct {
	Input int32
}

func ProcessConn(conn net.Conn, scene *Scene) {
	for {
		data := make([]byte, 1024)
		d := ServerResponse{}
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		data = []byte(strings.TrimSpace(string(data)))
		if err := json.Unmarshal(data[:n], &d); err != nil {
			continue
		}
		state := ClientSceneState{}
		for _, a := range d.Actors {
			aMap := a.(map[string]interface{})
			switch aMap["Type"] {
			case "player":
				dst := Player{}
				src := aMap["Actor"]
				jsonData, err := json.Marshal(src)
				if err != nil {
					panic(err)
				}
				if err = json.Unmarshal(jsonData, &dst); err != nil {
					panic(err)
				}
				state.Actors = append(state.Actors, &dst)
			case "ball":
				dst := Ball{}
				src := aMap["Actor"]
				jsonData, err := json.Marshal(src)
				if err != nil {
					panic(err)
				}
				if err = json.Unmarshal(jsonData, &dst); err != nil {
					panic(err)
				}
				state.Actors = append(state.Actors, &dst)
			default:
				log.Println("not supported")
			}
		}
		state.ApplyToScene(scene)
	}
}

type ServerResponse struct {
	Actors []interface{}
}
type ServerSceneState struct {
	Actors []ActorWrapper
}
type ActorWrapper struct {
	Type  string
	Actor Actor
}
type ClientSceneState struct {
	Actors []Actor
}

func (s *ClientSceneState) ApplyToScene(scene *Scene) {
	scene.Actors = s.Actors
}

func (g *Scene) GetSceneState() ServerSceneState {
	s := ServerSceneState{}
	for _, a := range g.Actors {
		s.Actors = append(s.Actors, a.ToActorWrapper())
	}
	return s
}
