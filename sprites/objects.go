package sprites

import (
	ut "main/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Objects struct {
	houses        [4]House
	trees         [6]Tree
	houseTexture  rl.Texture2D
	treeTexture   rl.Texture2D
	treeAnimation ut.Animation
	treeSourceRec rl.Rectangle
}

func (o *Objects) Load(houseTexture rl.Texture2D, treeTexture rl.Texture2D) {
	o.houseTexture = houseTexture
	o.treeTexture = treeTexture

	var x float32 = float32(ut.Globals.CanvasWidth)/2 - 64*15
	var y float32 = float32(ut.Globals.CanvasHeight)/2 - 64*15

	o.LoadHouses(x, y)
	o.LoadTrees(x, y)
}

func (o *Objects) LoadHouses(x float32, y float32) {
	o.houses[0] = House{Position: rl.NewVector2(x+300, y+300)}
	o.houses[1] = House{Position: rl.NewVector2(x+1400, y+200)}
	o.houses[2] = House{Position: rl.NewVector2(x+1000, y+1000)}
	o.houses[3] = House{Position: rl.NewVector2(x+500, y+1600)}
}

func (o *Objects) LoadTrees(x float32, y float32) {
	o.trees[0] = Tree{Position: rl.NewVector2(x+500, y+500)}
	o.trees[1] = Tree{Position: rl.NewVector2(x+900, y+200)}
	o.trees[2] = Tree{Position: rl.NewVector2(x+900, y+1500)}
	o.trees[3] = Tree{Position: rl.NewVector2(x+1500, y+500)}
	o.trees[4] = Tree{Position: rl.NewVector2(x+200, y+1200)}
	o.trees[5] = Tree{Position: rl.NewVector2(x+1550, y+1300)}

	o.treeAnimation = ut.Animation{
		FrameWidth:  uint16(o.treeTexture.Width) / 5,
		FrameHeight: uint16(o.treeTexture.Height),
		FrameX:      0,
		FrameY:      0,
		FrameCount:  5,
		FrameTime:   0.100,
	}
	o.treeSourceRec = rl.NewRectangle(0, 0, float32(o.treeAnimation.FrameWidth), float32(o.treeAnimation.FrameHeight))
}

func (o *Objects) Update() {
	o.UpdateTreeAnimation()
}

func (o *Objects) UpdateTreeAnimation() {
	var anim *ut.Animation = &o.treeAnimation
	anim.FrameTime -= rl.GetFrameTime()
	if anim.FrameTime < 0.0 {
		anim.FrameTime = 0.100
		anim.FrameX++
		if anim.FrameX == uint16(anim.FrameCount) {
			anim.FrameX = 0
		}
		o.treeSourceRec.X = float32(anim.FrameX) * float32(anim.FrameWidth)
	}
}

func (o *Objects) Draw() {
	for i := range o.houses {
		var house House = o.houses[i]
		rl.DrawTexture(o.houseTexture, int32(house.Position.X), int32(house.Position.Y), rl.White)
	}
	for i := range o.trees {
		var tree Tree = o.trees[i]
		rl.DrawTextureRec(o.treeTexture, o.treeSourceRec, tree.Position, rl.White)
	}
}

func (o *Objects) ApplyCameraOffset(offset rl.Vector2) {
	for i := range o.houses {
		var position *rl.Vector2 = &o.houses[i].Position
		position.X += offset.X
		position.Y += offset.Y
	}
	for i := range o.trees {
		var position *rl.Vector2 = &o.trees[i].Position
		position.X += offset.X
		position.Y += offset.Y
	}
}
