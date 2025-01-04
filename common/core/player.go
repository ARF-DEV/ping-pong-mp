package core

import (
	"encoding/json"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rect    rl.Rectangle
	UpKey   int32
	DownKey int32
	Color   rl.Color
}

func (p *Player) Update(scene *Scene) {
	inputMsg := InputMessage{}
	pressed := false
	dt := rl.GetFrameTime()
	if rl.IsKeyDown(p.UpKey) {
		p.Rect.Y -= PadSpeed * dt
		pressed = true
		inputMsg = InputMessage{
			Input: p.UpKey,
		}
	}
	if rl.IsKeyDown(p.DownKey) {
		p.Rect.Y += PadSpeed * dt
		pressed = true
		inputMsg = InputMessage{
			Input: p.DownKey,
		}
	}

	if pressed {
		data, _ := json.Marshal(inputMsg)
		scene.conn.Write(append(data, '\n'))
	}

}

func (p *Player) UpdateFromInput(in int32) {
	var dt float32 = float32(1) / 60
	if in == p.UpKey {
		p.Rect.Y -= PadSpeed * dt
	}
	if in == p.DownKey {
		p.Rect.Y += PadSpeed * dt
	}
}
func (p *Player) ToActorWrapper() ActorWrapper {
	return ActorWrapper{
		Type:  "player",
		Actor: p,
	}
}

func (p *Player) Draw() {
	pRect := p.Rect.ToInt32()
	rl.DrawRectangle(pRect.X, pRect.Y, pRect.Width, pRect.Height, p.Color)
}

func (p *Player) GetRect() rl.Rectangle {
	return p.Rect
}

func (p *Player) GetSnapShot() Player {
	return *p
}
