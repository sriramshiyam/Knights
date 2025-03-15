package sprites

import (
	"fmt"
	ut "main/utils"
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Object struct {
	Position     rl.Vector2
	Type         string
	CollisionBox ut.CollisionBox
}

type Objects struct {
	objects       [40]Object
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
	o.objects[0] = Object{Position: rl.NewVector2(x+300, y+300), Type: "house", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+300+7, y+300+33, 94, 69), AABBInfo: ut.AABBInfo{}}}
	o.objects[1] = Object{Position: rl.NewVector2(x+1400, y+200), Type: "house", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+1400+7, y+200+33, 94, 69), AABBInfo: ut.AABBInfo{}}}
	o.objects[2] = Object{Position: rl.NewVector2(x+1000, y+1000), Type: "house", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+1000+7, y+1000+33, 94, 69), AABBInfo: ut.AABBInfo{}}}
	o.objects[3] = Object{Position: rl.NewVector2(x+500, y+1600), Type: "house", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+500+7, y+1600+33, 94, 69), AABBInfo: ut.AABBInfo{}}}
}

func (o *Objects) LoadTrees(x float32, y float32) {
	o.objects[4] = Object{Position: rl.NewVector2(x+500, y+500), Type: "tree", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+500+88, y+500+120, 18, 10), AABBInfo: ut.AABBInfo{}}}
	o.objects[5] = Object{Position: rl.NewVector2(x+900, y+200), Type: "tree", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+900+88, y+200+120, 18, 10), AABBInfo: ut.AABBInfo{}}}
	o.objects[6] = Object{Position: rl.NewVector2(x+900, y+1500), Type: "tree", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+900+88, y+1500+120, 18, 10), AABBInfo: ut.AABBInfo{}}}
	o.objects[7] = Object{Position: rl.NewVector2(x+1500, y+500), Type: "tree", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+1500+88, y+500+120, 18, 10), AABBInfo: ut.AABBInfo{}}}
	o.objects[8] = Object{Position: rl.NewVector2(x+200, y+1200), Type: "tree", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+200+88, y+1200+120, 18, 10), AABBInfo: ut.AABBInfo{}}}
	o.objects[9] = Object{Position: rl.NewVector2(x+1550, y+1300), Type: "tree", CollisionBox: ut.CollisionBox{Rect: rl.NewRectangle(x+1550+88, y+1300+120, 18, 10), AABBInfo: ut.AABBInfo{}}}

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
	var plantIndex int8 = 10

	for _, tree := range o.objects[4:10] {
		o.objects[plantIndex] = Object{
			Position: rl.NewVector2(tree.Position.X+120+float32(rand.IntN(20)),
				tree.Position.Y+140+float32(rand.IntN(20))),
			Type: fmt.Sprintf("bush%d", spriteNumber)}

		plantIndex++
		spriteNumber++
		if spriteNumber > 3 {
			spriteNumber = 1
		}

		o.objects[plantIndex] = Object{
			Position: rl.NewVector2(tree.Position.X+float32(rand.IntN(20)),
				tree.Position.Y+110+float32(rand.IntN(20))),
			Type: fmt.Sprintf("bush%d", spriteNumber)}

		plantIndex++
		spriteNumber++

		if spriteNumber > 3 {
			spriteNumber = 1
		}
	}

	o.objects[22] = Object{Position: rl.NewVector2(x+100, y+100), Type: "grass1"}
	o.objects[23] = Object{Position: rl.NewVector2(x+300, y+500), Type: "grass2"}
	o.objects[24] = Object{Position: rl.NewVector2(x+700, y+700), Type: "mushroom1"}
	o.objects[25] = Object{Position: rl.NewVector2(x+100, y+1200), Type: "grass1"}
	o.objects[26] = Object{Position: rl.NewVector2(x+1100, y+900), Type: "grass2"}
	o.objects[27] = Object{Position: rl.NewVector2(x+1500, y+1500), Type: "mushroom2"}
	o.objects[28] = Object{Position: rl.NewVector2(x+800, y+1600), Type: "grass1"}
	o.objects[29] = Object{Position: rl.NewVector2(x+1200, y+1700), Type: "grass2"}
	o.objects[30] = Object{Position: rl.NewVector2(x+1600, y+200), Type: "mushroom1"}
	o.objects[31] = Object{Position: rl.NewVector2(x+1200, y+100), Type: "grass1"}
	o.objects[32] = Object{Position: rl.NewVector2(x+1500, y+800), Type: "grass2"}
	o.objects[33] = Object{Position: rl.NewVector2(x+900, y+780), Type: "mushroom2"}
	o.objects[34] = Object{Position: rl.NewVector2(x+1200, y+1500), Type: "grass1"}
	o.objects[35] = Object{Position: rl.NewVector2(x+200, y+1650), Type: "grass2"}
	o.objects[36] = Object{Position: rl.NewVector2(x+500, y+1200), Type: "mushroom1"}
	o.objects[37] = Object{Position: rl.NewVector2(x+550, y+1300), Type: "grass1"}
	o.objects[38] = Object{Position: rl.NewVector2(x+350, y+900), Type: "grass2"}
	o.objects[39] = Object{Position: rl.NewVector2(x+1200, y+1000), Type: "mushroom2"}
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
	for i := range o.objects {
		var object *Object = &o.objects[i]
		if object.Type == "house" {
			rl.DrawTexture(o.houseTexture, int32(object.Position.X), int32(object.Position.Y), rl.White)
			rl.DrawRectangleLinesEx(object.CollisionBox.Rect, 1, rl.Red)
		} else if object.Type == "tree" {
			rl.DrawTextureRec(o.treeTexture, o.treeSourceRec, object.Position, rl.White)
			rl.DrawRectangleLinesEx(object.CollisionBox.Rect, 1, rl.Red)
		} else {
			var texture rl.Texture2D = (*o.plantTextures)[object.Type]
			rl.DrawTexture(texture, int32(object.Position.X), int32(object.Position.Y), rl.White)
		}
	}
}

func ternary[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func (o *Objects) ApplyCameraOffset(offset rl.Vector2) {
	for i := range o.objects {
		var position *rl.Vector2 = &o.objects[i].Position
		position.X += offset.X
		position.Y += offset.Y
		var rect *rl.Rectangle = &o.objects[i].CollisionBox.Rect
		rect.X += offset.X
		rect.Y += offset.Y
	}
}
func (o *Objects) HandleCollisionWithKnight(knight *Knight) {
	var collisionBox *ut.CollisionBox = &knight.CollisionBox

	for i := range o.objects {
		var object *Object = &o.objects[i]
		var objectCollisionBox *ut.CollisionBox = &o.objects[i].CollisionBox
		var playerRect *rl.Rectangle = &collisionBox.Rect
		var objectRect *rl.Rectangle = &objectCollisionBox.Rect

		if (object.Type == "house" || object.Type == "tree") &&
			rl.Vector2Distance(rl.NewVector2(playerRect.X, playerRect.Y),
				rl.NewVector2(objectRect.X, objectRect.Y)) < 200 {

			var aabbInfo *ut.AABBInfo = &objectCollisionBox.AABBInfo

			var half1 rl.Vector2 = rl.NewVector2(playerRect.Width/2, playerRect.Height/2)
			var center1 rl.Vector2 = rl.NewVector2(playerRect.X+half1.X, playerRect.Y+half1.Y)

			var half2 rl.Vector2 = rl.NewVector2(objectRect.Width/2, objectRect.Height/2)
			var center2 rl.Vector2 = rl.NewVector2(objectRect.X+half2.X, objectRect.Y+half2.Y)

			var delta rl.Vector2 = rl.NewVector2(float32(math.Abs(float64(center1.X)-float64(center2.X))), float32(math.Abs(float64(center1.Y)-float64(center2.Y))))

			aabbInfo.OverlapX = half1.X + half2.X - delta.X
			aabbInfo.OverlapY = half1.Y + half2.Y - delta.Y

			if aabbInfo.OverlapX > 0 && aabbInfo.OverlapY > 0 {
				if aabbInfo.PreviousOverlapY > 0 {
					playerRect.X = playerRect.X + ternary(center1.X < center2.X, -aabbInfo.OverlapX, aabbInfo.OverlapX)
					knight.Position.X = playerRect.X - 74
				} else if aabbInfo.PreviousOverlapX > 0 {
					playerRect.Y = playerRect.Y + ternary(center1.Y < center2.Y, -aabbInfo.OverlapY, aabbInfo.OverlapY)
					knight.Position.Y = playerRect.Y - 68
				} else {
					playerRect.X = playerRect.X + ternary(center1.X < center2.X, -aabbInfo.OverlapX, aabbInfo.OverlapX)
					playerRect.Y = playerRect.Y + ternary(center1.Y < center2.Y, -aabbInfo.OverlapY, aabbInfo.OverlapY)
					knight.Position.X = playerRect.X - 74
					knight.Position.Y = playerRect.Y - 68
				}

			} else {
				aabbInfo.PreviousOverlapX = aabbInfo.OverlapX
				aabbInfo.PreviousOverlapY = aabbInfo.OverlapY
			}
		}
	}
}
