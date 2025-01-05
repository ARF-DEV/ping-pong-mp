package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Actor interface {
	Update(*Scene)
	ToActorWrapper() ActorWrapper
	Draw()
}

type PadActor interface {
	Actor
	UpdateFromInput(input int32)
	GetRect() rl.Rectangle
}
