package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Pos   rl.Vector2
	Rad   float32
	Color rl.Color
	Dir   rl.Vector2
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.Pos.X), int32(b.Pos.Y), b.Rad, b.Color)
}

func (b *Ball) Update(scene *Scene) {
	dt := rl.GetFrameTime()
	b.Pos = rl.Vector2Add(b.Pos, rl.Vector2Scale(b.Dir, BallSpeed*dt))
	areaTop := getAreaTop()
	if b.Pos.X+b.Rad >= areaTop.X+AreaWidth {
		b.Pos.X = areaTop.X + AreaWidth - b.Rad
		b.Dir.X *= -1
	}
	if b.Pos.X-b.Rad <= areaTop.X {
		b.Pos.X = areaTop.X + b.Rad
		b.Dir.X *= -1
	}
	if b.Pos.Y+b.Rad >= areaTop.Y+AreaHeight {
		b.Pos.Y = areaTop.Y + AreaHeight - b.Rad
		b.Dir.Y *= -1
	}
	if b.Pos.Y-b.Rad <= areaTop.Y {
		b.Pos.Y = areaTop.Y + b.Rad
		b.Dir.Y *= -1
	}

	for i := range scene.Actors {
		pad, ok := scene.Actors[i].(PadActor)
		if !ok {
			continue
		}
		padRect := pad.GetRect()
		if rl.CheckCollisionCircleRec(b.Pos, b.Rad, padRect) {
			padCenter := rl.Vector2{X: padRect.Width/2 + padRect.X, Y: padRect.Height/2 + padRect.Y}
			ballNewDir := rl.Vector2Normalize(rl.Vector2Subtract(b.Pos, padCenter))
			b.Dir = ballNewDir
		}
	}

}

func (b *Ball) UpdateFromInput(in int32) {}

func (b *Ball) ToActorWrapper() ActorWrapper {
	return ActorWrapper{
		Type:  "ball",
		Actor: b,
	}
}
func getAreaTop() rl.Vector2 {
	mWidth := rl.GetScreenWidth()
	mHeight := rl.GetScreenHeight()
	mCenterX := mWidth / 2
	mCenterY := mHeight / 2
	areaTopX := float32(mCenterX) - AreaWidth/2
	areaTopY := float32(mCenterY) - AreaHeight/2

	return rl.Vector2{
		X: areaTopX,
		Y: areaTopY,
	}
}
