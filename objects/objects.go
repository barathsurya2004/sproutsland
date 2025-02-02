package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Dest     rl.Rectangle
	Quantity int
}

func NewObject(
	dest rl.Rectangle,
) *Object {
	obj := Object{
		dest,
		3,
	}
	return &obj
}
