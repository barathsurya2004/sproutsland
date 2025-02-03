package class

import (
	"fmt"
	"strconv"

	"github.com/barathsurya2004/sproutsland/constants"
	"github.com/barathsurya2004/sproutsland/helpers"
	"github.com/barathsurya2004/sproutsland/objects"
	"github.com/barathsurya2004/sproutsland/scenes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Src           rl.Rectangle
	Dest          rl.Rectangle
	Tex           rl.Texture2D
	Speed         int
	playerFrame   int
	direction     int
	isMoving      bool
	Inventory     []objects.Object
	IsInteracting bool
}

func NewPlayer(url string) *Player {
	player := Player{}
	player.Dest = rl.NewRectangle(5*48, 15*48, 2*48, 2*48)
	player.Src = rl.NewRectangle(0, 0, 48, 48)
	player.Tex = rl.LoadTexture(url)
	player.Speed = 3
	return &player
}

func (p *Player) Move(gameFrame int, s *scenes.Scene) {
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
	if p.isMoving && !p.IsInteracting {
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

	dx, dy = p.isColliding(*s, dx, dy)

	if !p.IsInteracting {
		p.Dest.X += float32(dx)
		p.Dest.Y += float32(dy)
		p.Src.Y = p.Src.Height * float32(p.direction)
	}
	p.Src.X = p.Src.Width * float32(p.playerFrame)
}

func (p *Player) PickUpObject(s *scenes.Scene) {
	for i, object := range s.ObjectsPresent {
		if rl.CheckCollisionRecs(p.Dest, object.Dest) {
			p.Inventory = append(p.Inventory, object)
			s.ObjectsPresent = helpers.RemoveObjects(s.ObjectsPresent, i)
			fmt.Println("picking object ", i)
		}
	}
}

func (p *Player) isColliding(s scenes.Scene, dx, dy int) (int, int) {
	c := s.Collision
	for i, val := range c.Data {
		x := i % c.Width
		y := i / c.Width
		temp := rl.NewRectangle(float32(x)*constants.TileSize, float32(y)*constants.TileSize, constants.TileSize, constants.TileSize)
		cur := p.Dest
		cur.X += float32(dx)
		cur.Y += float32(dy)
		if rl.CheckCollisionRecs(cur, temp) && val != 0 {
			return 0, 0
		}
	}
	temp := p.Dest
	temp.X += float32(dx)
	temp.Y += float32(dy)
	if temp.X < 0 || temp.Y < 0 || temp.X+p.Dest.Width > float32(c.Width)*constants.TileSize || temp.Y+p.Dest.Height > float32(c.Height)*constants.TileSize {
		return 0, 0
	}
	return dx, dy
}

func (p *Player) Draw() {
	rl.DrawTexturePro(p.Tex, p.Src, p.Dest, rl.NewVector2(0, 0), 0, rl.White)
}

func (p *Player) DrawInventory() {
	if len(p.Inventory) == 0 {
		rl.DrawText("your Inventory is Empty", 48, 0, 32, rl.Black)
	}
	for i, object := range p.Inventory {
		temp := rl.NewRectangle(float32(i)*72, 0, 48, 48)
		rl.DrawTexturePro(object.Tex, rl.NewRectangle(0, 0, 16, 16), temp, rl.NewVector2(0, 0), 0, rl.White)
		quant := strconv.Itoa(object.Quantity)
		rl.DrawText(quant, int32(i*72+48), 48, 12, rl.Black)
	}
}
