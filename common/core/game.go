package core

import (
	"net"

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

func CreateGame() *Scene {
	mWidth := rl.GetScreenWidth()
	mHeight := rl.GetScreenHeight()
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

	// var err error
	// s.conn, err = net.Dial("tcp", ":8080")
	// if err != nil {
	// 	panic(err)
	// }

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
