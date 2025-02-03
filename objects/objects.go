package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Dest     rl.Rectangle
	Quantity int
	Tex      rl.Texture2D
}

func NewObject(
	dest rl.Rectangle,
	url string,
	quantity int,
) *Object {
	temp := rl.LoadTexture(url)
	obj := Object{
		dest,
		quantity,
		temp,
	}
	return &obj
}
