package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Actor interface {
	Update(*Scene)
	ToActorWrapper() ActorWrapper
	GetPos() rl.Vector2
	Draw()
}

type SnapShooter[T any] interface {
	GetSnapShot() T
}

type PadActor interface {
	Actor
	UpdateFromInput(input int32)
	GetRect() rl.Rectangle
}
