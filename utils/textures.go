package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	House  rl.Texture2D
	Knight rl.Texture2D
	Ground rl.Texture2D
	Foam   rl.Texture2D
	Tree   rl.Texture2D
}

func (t *Textures) Load() {
	t.Knight = rl.LoadTexture("res/image/knight.png")
	t.House = rl.LoadTexture("res/image/house.png")
	t.Ground = rl.LoadTexture("res/image/ground.png")
	t.Foam = rl.LoadTexture("res/image/foam.png")
	t.Tree = rl.LoadTexture("res/image/tree.png")
}

func (t *Textures) UnLoad() {
	rl.UnloadTexture(t.Knight)
	rl.UnloadTexture(t.House)
	rl.UnloadTexture(t.Ground)
	rl.UnloadTexture(t.Foam)
	rl.UnloadTexture(t.Tree)
}
