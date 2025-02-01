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
	Scenes    []*scenes.Scene
}

func NewGame() *Game {
	game := &Game{}
	game.P = class.NewPlayer("./assets/Characters/BasicCharSprite.png")
	game.Camera = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(game.P.Dest.X-screenWidth/2, game.P.Dest.Y-screenHeight/2), 0, 1)
	temp := scenes.NewScene("./assets/Tilesets/Grass.png", "./assets/TileMaps/tilees.tmj")
	game.Scenes = append(game.Scenes, temp)

	return game
}
