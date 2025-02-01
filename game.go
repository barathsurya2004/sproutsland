package main

import (
	class "github.com/barathsurya2004/sproutsland/Class"
	"github.com/barathsurya2004/sproutsland/scenes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	GameFrame int
	P         *class.Player
	Camera    rl.Camera2D
	Scene     *scenes.Scene
}

func NewGame() *Game {
	game := &Game{}
	game.P = class.NewPlayer("./assets/Characters/BasicCharSprite.png")
	game.Camera = rl.NewCamera2D(rl.NewVector2(screenWidth/2-game.P.Dest.Width/2, screenHeight/2-game.P.Dest.Height/2), rl.NewVector2(0, 0), 0, 1)
	game.Scene = scenes.NewScene("./assets/Tilesets/Grass.png", "./assets/TileMaps/tilees.tmj")

	return game
}
