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
	Src               rl.Rectangle
	Dest              rl.Rectangle
	Tex               rl.Texture2D
	Speed             int
	playerFrame       int
	direction         int
	isMoving          bool
	Inventory         []objects.Object
	IsInventoryOpen   bool
	IsInteracting     bool
	InventoryCS       int
	inventoryDropDown bool
	inventDropDownOp  int
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

	p.Src.X = p.Src.Width * float32(p.playerFrame)
	if !p.IsInteracting {
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
		dx, dy = p.isColliding(*s, dx, dy)
		p.Dest.X += float32(dx)
		p.Dest.Y += float32(dy)
		p.Src.Y = p.Src.Height * float32(p.direction)
	} else {
		if p.IsInventoryOpen {
			if !p.inventoryDropDown {
				if len(p.Inventory) != 0 {

					if rl.IsKeyPressed(rl.KeyEnter) {
						p.inventoryDropDown = true
					}
					if rl.IsKeyPressed(rl.KeyRight) {
						p.InventoryCS = (p.InventoryCS + 1) % len(p.Inventory)
					}
					if rl.IsKeyPressed(rl.KeyLeft) {
						p.InventoryCS = (p.InventoryCS - 1 + len(p.Inventory)) % len(p.Inventory)
					}
				}
			} else {
				sizeOfOp := len(p.Inventory[p.InventoryCS].Uses)
				if rl.IsKeyPressed(rl.KeyEnter) {
					if p.inventDropDownOp == 0 {
						p.UseObject()
					}
					p.inventoryDropDown = false
					p.inventDropDownOp = 0
				}
				if rl.IsKeyPressed(rl.KeyUp) {
					p.inventDropDownOp = (p.inventDropDownOp + 1 + sizeOfOp) % sizeOfOp
				}
				if rl.IsKeyPressed(rl.KeyDown) {
					p.inventDropDownOp = (p.inventDropDownOp - 1 + sizeOfOp) % sizeOfOp
				}
			}
		}
	}
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
		rl.DrawTexturePro(object.Tex, rl.NewRectangle(float32(object.Src)*16, 0, 16, 16), temp, rl.NewVector2(0, 0), 0, rl.White)
		quant := strconv.Itoa(object.Quantity)
		if p.InventoryCS == i {
			rl.DrawText(quant, int32(i*72+48), 48, 12, rl.Red)
			rl.DrawRectangleLinesEx(temp, 1, rl.Red)
		} else {
			rl.DrawText(quant, int32(i*72+48), 48, 12, rl.Black)
		}
		if p.inventoryDropDown && p.InventoryCS == i {
			rl.DrawRectangleV(rl.NewVector2(float32(i*72+24), 24), rl.NewVector2(48, 48), rl.Red)
			x := int32(i*72 + 24)
			for i, s := range p.Inventory[p.InventoryCS].Uses {
				if i == p.inventDropDownOp {
					rl.DrawText(s, x, int32(i*24+24), 15, rl.Blue)
				} else {
					rl.DrawText(s, x, int32(i*24+24), 15, rl.Black)
				}
			}
		}

	}
}

func (p *Player) UseObject() {
	if !p.Inventory[p.InventoryCS].CanReuse {
		p.Inventory = append(p.Inventory[:p.InventoryCS], p.Inventory[p.InventoryCS+1:]...)
	}
}
