package sprites

import (
	ut "main/utils"

	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type death struct {
	Position          rl.Vector2
	SourceRec         rl.Rectangle
	DestRec           rl.Rectangle
	Animation         ut.Animation
	AnimationFinished bool
}

type Deaths struct {
	list    []death
	texture rl.Texture2D
}

func (d *Deaths) Load(texture rl.Texture2D) {
	d.texture = texture
}

func (d *Deaths) UnLoad() {
	rl.UnloadTexture(d.texture)
}

func (d *Deaths) Update() {
	for i := len(d.list) - 1; i >= 0; i-- {
		var death *death = &d.list[i]
		var anim *ut.Animation = &death.Animation
		anim.FrameTime -= rl.GetFrameTime()
		if anim.FrameTime < 0.0 {
			anim.FrameTime = 0.100
			anim.FrameX++
			death.SourceRec.X = float32(anim.FrameX) * float32(anim.FrameWidth)
			if anim.FrameX == uint16(anim.FrameCount) {
				d.list = slices.Delete(d.list, i, i+1)
			}
		}
	}
	println(len(d.list))
}

func (d *Deaths) Draw() {
	for i := range d.list {
		rl.DrawTexturePro(d.texture, d.list[i].SourceRec, d.list[i].DestRec, rl.Vector2Zero(), 0, rl.White)
	}
}

func (d *Deaths) AddToList(position *rl.Vector2, invertSprite bool) {
	var death death = death{}
	death.Animation = ut.Animation{
		FrameWidth:  uint16(d.texture.Width / 14),
		FrameHeight: uint16(d.texture.Height),
		FrameX:      0,
		FrameY:      0,
		FrameCount:  14,
		FrameTime:   0.100}
	death.SourceRec = rl.NewRectangle(float32(death.Animation.FrameX), float32(death.Animation.FrameY), float32(death.Animation.FrameWidth), float32(death.Animation.FrameHeight))
	death.Position = rl.NewVector2(position.X-float32(death.Animation.FrameWidth)/2, position.Y-float32(death.Animation.FrameHeight)/2)
	death.DestRec = rl.NewRectangle(death.Position.X, death.Position.Y, float32(death.Animation.FrameWidth), float32(death.Animation.FrameHeight))
	if invertSprite {
		death.SourceRec.Width *= -1
	}
	d.list = append(d.list, death)
}

func (d *Deaths) ApplyCameraOffset(offset *rl.Vector2) {
	for i := range d.list {
		var position *rl.Vector2 = &d.list[i].Position
		position.X += offset.X
		position.Y += offset.Y
		var destRec *rl.Rectangle = &d.list[i].DestRec
		destRec.X += offset.X
		destRec.Y += offset.Y
	}
}
