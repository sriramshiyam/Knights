package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

type Ground struct {
	texture rl.Texture2D
	tiles   []Tile
}

type Tile struct {
	sourceRec rl.Rectangle
	position  rl.Vector2
}

func (g *Ground) LoadTiles() {
	g.tiles = make([]Tile, 0, 400)
	var tilesNumber [20][20]int

	tilesNumber[0][0] = 1
	for i := 1; i <= 18; i++ {
		tilesNumber[0][i] = 2
	}
	tilesNumber[0][19] = 3

	for i := 1; i <= 18; i++ {
		tilesNumber[i][0] = 4
		for j := 1; j <= 18; j++ {
			tilesNumber[i][j] = 5
		}
		tilesNumber[i][19] = 6
	}

	tilesNumber[19][0] = 7
	for i := 1; i <= 18; i++ {
		tilesNumber[19][i] = 8
	}
	tilesNumber[19][19] = 9

	var x float32 = 0
	var y float32 = 0

	for i := range tilesNumber {
		for j := range tilesNumber[i] {
			var tile int = tilesNumber[i][j]
			var sourceRec rl.Rectangle = rl.Rectangle{}
			sourceRec.Width = 64
			sourceRec.Height = 64
			for tile > 3 {
				tile -= 3
				sourceRec.Y += 64
			}
			sourceRec.X = float32(tile-1) * 64
			g.tiles = append(g.tiles, Tile{sourceRec, rl.NewVector2(x, y)})
			x += 64
		}
		x = 0
		y += 64
	}
}

func (g *Ground) Load() {
	g.texture = rl.LoadTexture("res/image/ground.png")
	g.LoadTiles()

}

func (g *Ground) UnLoad() {
	rl.UnloadTexture(g.texture)
}

func (g *Ground) Update() {

}

func (g *Ground) Draw() {
	for i := range g.tiles {
		rl.DrawTextureRec(g.texture, g.tiles[i].sourceRec, g.tiles[i].position, rl.White)
	}
}

func (g *Ground) ApplyCameraOffset(offset rl.Vector2) {
	for i := range g.tiles {
		var tilePosition *rl.Vector2 = &g.tiles[i].position
		tilePosition.X += offset.X
		tilePosition.Y += offset.Y
	}
}
