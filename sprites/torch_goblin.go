package sprites

import (
	ut "main/utils"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TorchGoblin struct {
	Position     rl.Vector2
	direction    rl.Vector2
	origin       rl.Vector2
	texture      rl.Texture2D
	animation    ut.Animation
	sourceRec    rl.Rectangle
	destRec      rl.Rectangle
	state        string
	Speed        float32
	CollisionBox ut.CollisionBox
	aimSide      int
	player       *Knight
	moving       bool
}

func (t *TorchGoblin) Load(texture rl.Texture2D, player *Knight) {
	t.player = player
	t.state = "idle"
	t.texture = texture
	t.Speed = 300
	t.animation = ut.Animation{
		FrameX:      0,
		FrameCount:  7,
		FrameWidth:  uint16(t.texture.Width) / 7,
		FrameHeight: uint16(t.texture.Height) / 5,
		FrameY:      0,
		FrameTime:   0.100,
	}
	t.Position = rl.NewVector2(float32(ut.Globals.CanvasWidth)/2-float32(t.animation.FrameWidth)/2-300, float32(ut.Globals.CanvasHeight)/2-float32(t.animation.FrameHeight)/2)
	t.sourceRec = rl.Rectangle{
		X:      float32(t.animation.FrameX),
		Y:      float32(t.animation.FrameY),
		Width:  float32(t.animation.FrameWidth),
		Height: float32(t.animation.FrameHeight),
	}
	t.destRec = rl.Rectangle{
		X:      float32(t.animation.FrameX),
		Y:      float32(t.animation.FrameY),
		Width:  float32(t.animation.FrameWidth),
		Height: float32(t.animation.FrameHeight),
	}
	t.destRec.X = t.Position.X
	t.destRec.Y = t.Position.Y
	t.origin = rl.Vector2{}
	t.CollisionBox = ut.CollisionBox{Rect: rl.NewRectangle(t.Position.X+77, t.Position.Y+65, 37, 63)}
}

func (t *TorchGoblin) UnLoad() {

}

func (t *TorchGoblin) Update() {
	var playerRect *rl.Rectangle = &t.player.CollisionBox.Rect
	var playerPos rl.Vector2 = rl.NewVector2(playerRect.X+playerRect.Width/2, playerRect.Y+playerRect.Height/2)

	var rect *rl.Rectangle = &t.CollisionBox.Rect
	var pos rl.Vector2 = rl.NewVector2(rect.X+rect.Width/2, rect.Y+rect.Height/2)

	if playerPos.X < pos.X {
		if t.sourceRec.Width > 0 {
			t.sourceRec.Width *= -1
		}
	} else {
		if t.sourceRec.Width < 0 {
			t.sourceRec.Width *= -1
		}
	}

	t.direction = rl.Vector2Normalize(rl.Vector2Subtract(playerPos, pos))
	var distance float32 = rl.Vector2Distance(playerPos, pos)

	if distance > 70 {
		t.Position = rl.Vector2Add(t.Position, rl.Vector2Scale(t.direction, t.Speed*rl.GetFrameTime()))
		t.destRec.X = t.Position.X
		t.destRec.Y = t.Position.Y
		t.CollisionBox.Rect.X = t.Position.X + 77
		t.CollisionBox.Rect.Y = t.Position.Y + 65
		t.moving = true
	} else {
		t.moving = false
	}

	if t.moving && t.state != "running" {
		t.state = "running"
		t.UpdateState()
	} else if !t.moving && t.state != "idle" {
		t.state = "idle"
		t.UpdateState()
	}

	t.UpdateAnimation()
}

func (t *TorchGoblin) UpdateAnimation() {
	var anim *ut.Animation = &t.animation
	anim.FrameTime -= rl.GetFrameTime()
	if anim.FrameTime < 0.0 {
		anim.FrameTime = 0.100
		anim.FrameX++
		if anim.FrameX == uint16(anim.FrameCount) {
			anim.FrameX = 0
			if strings.Contains(t.state, "attack") {
				t.state = "idle"
				t.UpdateState()
			}
		}
		t.sourceRec.X = float32(anim.FrameX) * float32(anim.FrameWidth)
	}
}

func (t *TorchGoblin) UpdateState() {
	var anim *ut.Animation = &t.animation
	switch t.state {
	case "idle":
		anim.FrameCount = 7
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = 0
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	case "running":
		anim.FrameCount = 6
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	}
}

func (t *TorchGoblin) Draw() {
	rl.DrawTexturePro(t.texture, t.sourceRec, t.destRec, t.origin, 0, rl.White)
	// rl.DrawRectangleLinesEx(t.CollisionBox.Rect, 1, rl.Red)
}

func (t *TorchGoblin) ApplyCameraOffset(offset rl.Vector2) {
	t.Position.X += offset.X
	t.Position.Y += offset.Y
	t.destRec.X += offset.X
	t.destRec.Y += offset.Y
	t.CollisionBox.Rect.X += offset.X
	t.CollisionBox.Rect.Y += offset.Y
}
