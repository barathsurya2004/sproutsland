package scenes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/barathsurya2004/sproutsland/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	TileSetTexs    []rl.Texture2D
	TileSetJson    *tileSetJson
	texHeight      int
	texWidth       int
	Src            rl.Rectangle
	Dest           rl.Rectangle
	Collision      TilesLayer
	ObjectsPresent []objects.Object
}

type tileSetJson struct {
	Layers []*TilesLayer `json:"layers"`
}

type TilesLayer struct {
	Data     []int  `json:"data"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	FileDest string `json:"url"`
	Name     string `json:"name"`
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
	for j, layer := range s.TileSetJson.Layers {
		if layer.Name == "collision" {
			continue
		}
		for i, tile := range layer.Data {
			x := (tile - 1) % s.texWidth
			y := (tile - 1) / s.texWidth
			s.Src.X = s.Src.Width * float32(x)
			s.Src.Y = s.Src.Width * float32(y)
			s.Dest.X = float32((i % layer.Width) * 48)
			s.Dest.Y = float32((i / layer.Width) * 48)
			rl.DrawTexturePro(s.TileSetTexs[j], s.Src, s.Dest, rl.NewVector2(0, 0), 0, rl.White)
		}
	}
	for _, val := range s.ObjectsPresent {
		x := val.Src
		src := rl.NewRectangle(float32(x)*16, 0, 16, 16)
		rl.DrawTexturePro(val.Tex, src, val.Dest, rl.NewVector2(0, 0), 0, rl.White)
	}
}

func NewScene(tilejsonUrl string) *Scene {
	temp := &Scene{}

	temp.deconstructJson(tilejsonUrl)
	temp.texHeight = 5
	temp.texWidth = 11
	temp.Src = rl.NewRectangle(0, 0, 16, 16)
	temp.Dest = rl.NewRectangle(0, 0, 48, 48)
	for i := range temp.TileSetJson.Layers {
		if temp.TileSetJson.Layers[i].Name != "collision" {
			temp.TileSetTexs = append(temp.TileSetTexs, rl.LoadTexture(temp.TileSetJson.Layers[i].FileDest))
		} else {
			temp.Collision = *temp.TileSetJson.Layers[i]
		}
		fmt.Println(temp.TileSetTexs)
	}

	temp.ObjectsPresent = append(temp.ObjectsPresent, *objects.NewObject(rl.NewRectangle(0, 0, 48, 48), "./assets/Objects/Simple_Milk_and_grass_item.png", 2))
	temp.ObjectsPresent = append(temp.ObjectsPresent, *objects.NewObject(rl.NewRectangle(15*48, 10*48, 48, 48), "./assets/Objects/Simple_Milk_and_grass_item.png", 1))
	temp.ObjectsPresent = append(temp.ObjectsPresent, *objects.NewObject(rl.NewRectangle(20*48, 20*48, 48, 48), "./assets/Objects/Simple_Milk_and_grass_item.png", 5))
	temp.ObjectsPresent = append(temp.ObjectsPresent, *objects.NewObject(rl.NewRectangle(0, 15*48, 48, 48), "./assets/Objects/Simple_Milk_and_grass_item.png", 3))
	return temp
}
