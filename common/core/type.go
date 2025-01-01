package core

import rl "github.com/gen2brain/raylib-go/raylib"

type Actor interface {
	Update(*Scene)
	Draw()
}

type PadActor interface {
	Actor
	GetRect() rl.Rectangle
}
