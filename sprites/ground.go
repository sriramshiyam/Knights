package sprites

import (
	ut "main/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ground struct {
	tileTexture   rl.Texture2D
	foamTexture   rl.Texture2D
	foamAnimation ut.Animation
	foamSourceRec rl.Rectangle
	foamOrigin    rl.Vector2
	tiles         []Tile
}

type Tile struct {
	sourceRec    rl.Rectangle
	position     rl.Vector2
	isCornerTile bool
}

func (g *Ground) LoadTiles() {
	g.tiles = make([]Tile, 0, 900)
	var tilesNumber [30][30]int

	tilesNumber[0][0] = 1
	for i := 1; i <= 29; i++ {
		tilesNumber[0][i] = 2
	}
	tilesNumber[0][29] = 3

	for i := 1; i <= 29; i++ {
		tilesNumber[i][0] = 4
		for j := 1; j <= 29; j++ {
			tilesNumber[i][j] = 5
		}
		tilesNumber[i][29] = 6
	}

	tilesNumber[29][0] = 7
	for i := 1; i <= 29; i++ {
		tilesNumber[29][i] = 8
	}
	tilesNumber[29][29] = 9

	var x float32 = float32(ut.Globals.CanvasWidth)/2 - 64*15
	var y float32 = float32(ut.Globals.CanvasHeight)/2 - 64*15

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
			g.tiles = append(g.tiles, Tile{sourceRec: sourceRec, position: rl.NewVector2(x, y), isCornerTile: (i == 0 || j == 0 || i == 29 || j == 29)})
			x += 64
		}
		x = float32(ut.Globals.CanvasWidth)/2 - 64*15
		y += 64
	}
}

func (g *Ground) Load(tileTexture rl.Texture2D, foamTexture rl.Texture2D) {
	g.tileTexture = tileTexture
	g.foamTexture = foamTexture
	g.foamAnimation = ut.Animation{
		FrameCount:  8,
		FrameWidth:  uint16(g.foamTexture.Width) / 8,
		FrameHeight: uint16(g.foamTexture.Height),
		FrameX:      0,
		FrameY:      0,
		FrameTime:   0.100,
	}
	g.foamSourceRec = rl.NewRectangle(0, 0, 192, 192)
	g.foamOrigin = rl.NewVector2(64, 64)
	g.LoadTiles()
}

func (g *Ground) UnLoad() {
	rl.UnloadTexture(g.tileTexture)
	rl.UnloadTexture(g.foamTexture)
}

func (g *Ground) Update() {
	g.UpdateFoamAnimation()
}

func (g *Ground) UpdateFoamAnimation() {
	var anim *ut.Animation = &g.foamAnimation
	anim.FrameTime -= rl.GetFrameTime()
	if anim.FrameTime < 0.0 {
		anim.FrameTime = 0.100
		anim.FrameX++
		if anim.FrameX == uint16(anim.FrameCount) {
			anim.FrameX = 0
		}
		g.foamSourceRec.X = float32(anim.FrameX * anim.FrameWidth)
	}
}

func (g *Ground) Draw() {
	for i := range g.tiles {
		if g.tiles[i].isCornerTile {
			var position *rl.Vector2 = &g.tiles[i].position
			rl.DrawTexturePro(g.foamTexture, g.foamSourceRec, rl.NewRectangle(position.X, position.Y, g.foamSourceRec.Width, g.foamSourceRec.Height), g.foamOrigin, 0, rl.White)
		}
	}
	for i := range g.tiles {
		rl.DrawTextureRec(g.tileTexture, g.tiles[i].sourceRec, g.tiles[i].position, rl.White)
	}
}

func (g *Ground) ApplyCameraOffset(offset rl.Vector2) {
	for i := range g.tiles {
		var tilePosition *rl.Vector2 = &g.tiles[i].position
		tilePosition.X += offset.X
		tilePosition.Y += offset.Y
	}
}
