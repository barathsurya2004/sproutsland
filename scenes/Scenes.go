package scenes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	TileSetTex  rl.Texture2D
	TileSetJson *tileSetJson
}

type tileSetJson struct {
	Layers *tilesLayer `json:"layers"`
}

type tilesLayer struct {
	Data []int `json:"data"`
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
	fmt.Println(s.TileSetJson)
}

func NewScene(tilesetUrl, tilejsonUrl string) *Scene {
	temp := &Scene{}

	temp.TileSetTex = rl.LoadTexture(tilesetUrl)

	temp.deconstructJson(tilejsonUrl)

	return temp
}
