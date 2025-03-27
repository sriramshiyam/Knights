package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	House         rl.Texture2D
	Knight        rl.Texture2D
	Ground        rl.Texture2D
	Foam          rl.Texture2D
	Tree          rl.Texture2D
	PlantTextures map[string]rl.Texture2D
	TorchGolbin   rl.Texture2D
	Death         rl.Texture2D
}

func (t *Textures) Load() {
	t.Knight = rl.LoadTexture("res/image/knight.png")
	t.House = rl.LoadTexture("res/image/house.png")
	t.Ground = rl.LoadTexture("res/image/ground.png")
	t.Foam = rl.LoadTexture("res/image/foam.png")
	t.Tree = rl.LoadTexture("res/image/tree.png")
	t.PlantTextures = make(map[string]rl.Texture2D)
	t.PlantTextures["bush1"] = rl.LoadTexture("res/image/bush1.png")
	t.PlantTextures["bush2"] = rl.LoadTexture("res/image/bush2.png")
	t.PlantTextures["bush3"] = rl.LoadTexture("res/image/bush3.png")
	t.PlantTextures["grass1"] = rl.LoadTexture("res/image/grass1.png")
	t.PlantTextures["grass2"] = rl.LoadTexture("res/image/grass2.png")
	t.PlantTextures["mushroom1"] = rl.LoadTexture("res/image/mushroom1.png")
	t.PlantTextures["mushroom2"] = rl.LoadTexture("res/image/mushroom2.png")
	t.TorchGolbin = rl.LoadTexture("res/image/torchgoblin.png")
	t.Death = rl.LoadTexture("res/image/dead.png")
}

func (t *Textures) UnLoad() {
	rl.UnloadTexture(t.Knight)
	rl.UnloadTexture(t.House)
	rl.UnloadTexture(t.Ground)
	rl.UnloadTexture(t.Foam)
	rl.UnloadTexture(t.Tree)
	for _, texture := range t.PlantTextures {
		rl.UnloadTexture(texture)
	}
	rl.UnloadTexture(t.TorchGolbin)
	rl.UnloadTexture(t.Death)
}
