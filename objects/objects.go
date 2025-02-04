package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Dest     rl.Rectangle
	Quantity int
	Tex      rl.Texture2D
	Src      int32
	Uses     []string
	CanReuse bool
}

func NewObject(
	dest rl.Rectangle,
	url string,
	quantity int,
) *Object {
	temp := rl.LoadTexture(url)
	x := rl.GetRandomValue(0, 3)
	obj := Object{
		dest,
		quantity,
		temp,
		x,
		[]string{
			"Drink",
			"Throw",
		},
		false,
	}
	return &obj
}
