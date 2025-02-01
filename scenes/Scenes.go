package scenes

import (
	"encoding/json"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	TileSetTex  rl.Texture2D
	TileSetJson *tileSetJson
	texHeight   int
	texWidth    int
	Src         rl.Rectangle
	Dest        rl.Rectangle
}

type tileSetJson struct {
	Layers []*tilesLayer `json:"layers"`
}

type tilesLayer struct {
	Data   []int `json:"data"`
	Height int   `json:"height"`
	Width  int   `json:"width"`
}

func (s *Scene) deconstructJson(url string) {
	file, err := os.ReadFile(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	var tilejson tileSetJson
	err = json.Unmarshal(file, &tilejson)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	s.TileSetJson = &tilejson
}

func (s *Scene) DrawScene() {
	for _, layer := range s.TileSetJson.Layers {
		for i, tile := range layer.Data {
			x := (tile - 1) % s.texWidth
			y := (tile - 1) / s.texWidth
			s.Src.X = s.Src.Width * float32(x)
			s.Src.Y = s.Src.Width * float32(y)
			s.Dest.X = float32((i % layer.Width) * 48)
			s.Dest.Y = float32((i / layer.Width) * 48)
			rl.DrawTexturePro(s.TileSetTex, s.Src, s.Dest, rl.NewVector2(0, 0), 0, rl.White)

		}
	}
}

func NewScene(tilesetUrl, tilejsonUrl string) *Scene {
	temp := &Scene{}

	temp.TileSetTex = rl.LoadTexture(tilesetUrl)

	temp.deconstructJson(tilejsonUrl)
	temp.texHeight = 5
	temp.texWidth = 11
	temp.Src = rl.NewRectangle(0, 0, 16, 16)
	temp.Dest = rl.NewRectangle(0, 0, 48, 48)

	return temp
}
