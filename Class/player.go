package class

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Src         rl.Rectangle
	Dest        rl.Rectangle
	Tex         rl.Texture2D
	Speed       int
	playerFrame int
	direction   int
	isMoving    bool
}

func NewPlayer(url string) *Player {
	player := Player{}
	player.Dest = rl.NewRectangle(0, 0, 100, 100)
	player.Src = rl.NewRectangle(0, 0, 48, 48)
	player.Tex = rl.LoadTexture(url)
	player.Speed = 3
	return &player
}

func (p *Player) Move(gameFrame int) {
	dx, dy := 0, 0
	p.isMoving = false
	if rl.IsKeyDown(rl.KeyUp) {
		dy -= p.Speed
		p.direction = 1
		p.isMoving = true
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		dx -= p.Speed
		p.direction = 2
		p.isMoving = true
	}
	if rl.IsKeyDown(rl.KeyRight) {
		dx += p.Speed
		p.direction = 3
		p.isMoving = true
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.direction = 0
		dy += p.Speed
		p.isMoving = true
	}
	if p.isMoving {
		if gameFrame%10 == 1 {
			p.playerFrame += 1
		}

		if p.playerFrame >= 4 {
			p.playerFrame = 0
		}
	} else {
		if gameFrame%20 == 1 {
			p.playerFrame += 1
		}

		if p.playerFrame >= 2 {
			p.playerFrame = 0
		}
	}

	p.Dest.X += float32(dx)
	p.Dest.Y += float32(dy)
	p.Src.X = p.Src.Width * float32(p.playerFrame)
	p.Src.Y = p.Src.Height * float32(p.direction)
}

func (p *Player) Draw() {
	rl.DrawTexturePro(p.Tex, p.Src, p.Dest, rl.NewVector2(0, 0), 0, rl.White)
}
