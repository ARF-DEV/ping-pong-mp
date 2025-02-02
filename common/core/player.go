package core

import (
	"github.com/ARF-DEV/ping-pong-mp/common/network"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rect    rl.Rectangle
	UpKey   int32
	DownKey int32
	Color   rl.Color
}

func (p *Player) Update(scene *Scene) {
	inputMsg := network.InputMessage{}
	pressed := false
	dt := float32(1) / network.SERVER_TICK
	if rl.IsKeyDown(p.UpKey) {
		p.Rect.Y -= PadSpeed * dt
		pressed = true
		inputMsg = network.InputMessage{
			Tick:  scene.tick,
			Input: p.UpKey,
		}
	}
	if rl.IsKeyDown(p.DownKey) {
		p.Rect.Y += PadSpeed * dt
		pressed = true
		inputMsg = network.InputMessage{
			Tick:  scene.tick,
			Input: p.DownKey,
		}
	}

	if pressed {
		scene.curPressedKey = append(scene.curPressedKey, inputMsg.Input)
		network.Send(scene.conn, inputMsg)
	}

}

func (p *Player) GetSnapShot() Player {
	return *p
}

func (p *Player) UpdateFromInput(in int32) {
	var dt float32 = float32(1) / network.SERVER_TICK
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
func (p *Player) GetPos() rl.Vector2 {
	return rl.Vector2{X: p.Rect.X, Y: p.Rect.Y}
}

func (p *Player) Draw() {
	pRect := p.Rect.ToInt32()
	rl.DrawRectangle(pRect.X, pRect.Y, pRect.Width, pRect.Height, p.Color)
}

func (p *Player) GetRect() rl.Rectangle {
	return p.Rect
}
