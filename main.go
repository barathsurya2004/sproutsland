package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenHeight = 720
	screenWidth  = 1240
)

var (
	game  *Game
	music rl.Music
)

func InitGame() {
	rl.InitWindow(screenWidth, screenHeight, "SproutsLand")
	game = NewGame()
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("./assets/bgm.mp3")
	rl.PlayMusicStream(music)
}

func QuitGame() {
	rl.UnloadMusicStream(music)
	rl.UnloadTexture(game.P.Tex)
	rl.CloseWindow()
}

func Input() {
	game.P.Move(game.GameFrame, game.Scenes[0])
}

func Update() {
	game.GameFrame += 1
	dx := math.Max(float64(game.P.Dest.X)-screenWidth/2, 0)
	dy := math.Max(float64(game.P.Dest.Y)-screenHeight/2, 0)
	dx = math.Min(dx, float64(game.Scenes[0].TileSetJson.Layers[0].Width*48-screenWidth))
	dy = math.Min(dy, float64(game.Scenes[0].TileSetJson.Layers[0].Height*48-screenHeight))
	game.Camera.Target = rl.NewVector2(float32(dx), float32(dy))
	rl.UpdateMusicStream(music)
}

func Draw() {
	rl.BeginDrawing()
	rl.BeginMode2D(game.Camera)
	rl.ClearBackground(rl.RayWhite)
	game.Scenes[0].DrawScene()
	game.P.Draw()
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
