package sprites

import (
	"fmt"
	ut "main/utils"
	"math/rand/v2"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Objects struct {
	houses        [4]House
	trees         [6]Tree
	plants        [30]Plant
	houseTexture  rl.Texture2D
	treeTexture   rl.Texture2D
	treeAnimation ut.Animation
	treeSourceRec rl.Rectangle
	plantTextures *map[string]rl.Texture2D
}

func (o *Objects) Load(houseTexture rl.Texture2D, treeTexture rl.Texture2D, plantTextures *map[string]rl.Texture2D) {
	o.houseTexture = houseTexture
	o.treeTexture = treeTexture
	o.plantTextures = plantTextures

	var x float32 = float32(ut.Globals.CanvasWidth)/2 - 64*15
	var y float32 = float32(ut.Globals.CanvasHeight)/2 - 64*15

	o.LoadHouses(x, y)
	o.LoadTrees(x, y)
	o.LoadPlants(x, y)
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

func (o *Objects) LoadPlants(x float32, y float32) {
	var spriteNumber int8 = 1
	var plantIndex int8 = 0

	for _, tree := range o.trees {
		o.plants[plantIndex] = Plant{
			Position: rl.NewVector2(tree.Position.X+120+float32(rand.IntN(20)),
				tree.Position.Y+140+float32(rand.IntN(20))),
			PlantType: fmt.Sprintf("bush%d", spriteNumber)}

		plantIndex++
		spriteNumber++
		if spriteNumber > 3 {
			spriteNumber = 1
		}

		o.plants[plantIndex] = Plant{
			Position: rl.NewVector2(tree.Position.X+float32(rand.IntN(20)),
				tree.Position.Y+110+float32(rand.IntN(20))),
			PlantType: fmt.Sprintf("bush%d", spriteNumber)}

		plantIndex++
		spriteNumber++

		if spriteNumber > 3 {
			spriteNumber = 1
		}
	}

	o.plants[12] = Plant{Position: rl.NewVector2(x+100, y+100), PlantType: "grass1"}
	o.plants[13] = Plant{Position: rl.NewVector2(x+300, y+500), PlantType: "grass2"}
	o.plants[14] = Plant{Position: rl.NewVector2(x+700, y+700), PlantType: "mushroom1"}
	o.plants[15] = Plant{Position: rl.NewVector2(x+100, y+1200), PlantType: "grass1"}
	o.plants[16] = Plant{Position: rl.NewVector2(x+1100, y+900), PlantType: "grass2"}
	o.plants[17] = Plant{Position: rl.NewVector2(x+1500, y+1500), PlantType: "mushroom2"}
	o.plants[18] = Plant{Position: rl.NewVector2(x+800, y+1600), PlantType: "grass1"}
	o.plants[19] = Plant{Position: rl.NewVector2(x+1200, y+1700), PlantType: "grass2"}
	o.plants[20] = Plant{Position: rl.NewVector2(x+1600, y+200), PlantType: "mushroom1"}
	o.plants[21] = Plant{Position: rl.NewVector2(x+1200, y+100), PlantType: "grass1"}
	o.plants[22] = Plant{Position: rl.NewVector2(x+1500, y+800), PlantType: "grass2"}
	o.plants[23] = Plant{Position: rl.NewVector2(x+900, y+780), PlantType: "mushroom2"}
	o.plants[24] = Plant{Position: rl.NewVector2(x+1200, y+1500), PlantType: "grass1"}
	o.plants[25] = Plant{Position: rl.NewVector2(x+200, y+1650), PlantType: "grass2"}
	o.plants[26] = Plant{Position: rl.NewVector2(x+500, y+1200), PlantType: "mushroom1"}
	o.plants[27] = Plant{Position: rl.NewVector2(x+550, y+1300), PlantType: "grass1"}
	o.plants[28] = Plant{Position: rl.NewVector2(x+350, y+900), PlantType: "grass2"}
	o.plants[29] = Plant{Position: rl.NewVector2(x+1200, y+1000), PlantType: "mushroom2"}
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
	for i := range o.plants {
		var plant Plant = o.plants[i]
		var texture rl.Texture2D = (*o.plantTextures)[plant.PlantType]
		rl.DrawTexture(texture, int32(plant.Position.X), int32(plant.Position.Y), rl.White)
		if strings.Contains(plant.PlantType, "grass") || strings.Contains(plant.PlantType, "mushroom") {
			rl.DrawRectangle(int32(plant.Position.X), int32(plant.Position.Y), texture.Width, texture.Height, rl.Red)
			rl.DrawText(fmt.Sprintf("%d %s", i, plant.PlantType), int32(plant.Position.X), int32(plant.Position.Y), 20, rl.White)
		}
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
	for i := range o.plants {
		var position *rl.Vector2 = &o.plants[i].Position
		position.X += offset.X
		position.Y += offset.Y
	}
}
