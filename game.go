package main

import (
	class "github.com/barathsurya2004/sproutsland/Class"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	GameFrame int
	P         *class.Player
	Camera    rl.Camera2D
}
