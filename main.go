package main

import (
	class "github.com/barathsurya2004/sproutsland/Class"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenHeight = 720
	screenWidth  = 1240
)

var game *Game

func InitGame() {
	rl.InitWindow(screenWidth, screenHeight, "SproutsLand")
	game = &Game{}
	game.P = class.NewPlayer("./assets/Characters/BasicCharSprite.png")
	game.Camera = rl.NewCamera2D(rl.NewVector2(screenWidth/2-game.P.Dest.Width/2, screenHeight/2-game.P.Dest.Height/2), rl.NewVector2(0, 0), 0, 1)

	rl.SetTargetFPS(60)
}

func QuitGame() {
	rl.UnloadTexture(game.P.Tex)
	rl.CloseWindow()
}

func Input() {
	game.P.Move(game.GameFrame)
}

func Update() {
	game.GameFrame += 1
	game.Camera.Target = rl.NewVector2(game.P.Dest.X, game.P.Dest.Y)
}

func Draw() {
	rl.BeginDrawing()
	rl.BeginMode2D(game.Camera)
	rl.ClearBackground(rl.RayWhite)
	game.P.Draw()
	rl.DrawText("heeloso", 0, 0, 55, rl.Red)
	rl.EndMode2D()

	rl.EndDrawing()
}

func main() {
	InitGame()
	defer QuitGame()
	for !rl.WindowShouldClose() {
		Input()
		Update()
		Draw()
	}
}
