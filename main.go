package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenHeight = 720
	screenWidth  = 1240
)

var game *Game

func InitGame() {
	rl.InitWindow(screenWidth, screenHeight, "SproutsLand")
	game = NewGame()
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
