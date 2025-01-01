package core

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Rect    rl.Rectangle
	UpKey   int32
	DownKey int32
	Color   rl.Color
}

func (p *Player) Update(scene *Scene) {
	dt := rl.GetFrameTime()
	if rl.IsKeyDown(p.UpKey) {
		p.Rect.Y -= PadSpeed * dt
	}
	if rl.IsKeyDown(p.DownKey) {
		p.Rect.Y += PadSpeed * dt
	}

}

func (p *Player) Draw() {
	pRect := p.Rect.ToInt32()
	rl.DrawRectangle(pRect.X, pRect.Y, pRect.Width, pRect.Height, p.Color)
}

func (p *Player) GetRect() rl.Rectangle {
	return p.Rect
}
