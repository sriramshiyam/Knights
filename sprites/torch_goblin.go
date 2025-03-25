package sprites

import (
	ut "main/utils"
	"math"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TorchGoblin struct {
	Position          rl.Vector2
	direction         rl.Vector2
	origin            rl.Vector2
	texture           rl.Texture2D
	animation         ut.Animation
	sourceRec         rl.Rectangle
	destRec           rl.Rectangle
	state             string
	Speed             float32
	CollisionBox      ut.CollisionBox
	aimSide           int
	player            *Knight
	moving            bool
	attacked          bool
	attackedForce     float32
	attackedDirection rl.Vector2
	waitTimer         float32
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

	if !t.attacked && distance > 70 && !strings.Contains(t.state, "attack") {
		t.Position = rl.Vector2Add(t.Position, rl.Vector2Scale(t.direction, t.Speed*rl.GetFrameTime()))
		t.destRec.X = t.Position.X
		t.destRec.Y = t.Position.Y
		t.CollisionBox.Rect.X = t.Position.X + 77
		t.CollisionBox.Rect.Y = t.Position.Y + 65
		t.moving = true
	} else {
		t.moving = false
		if !t.attacked {
			t.handleAimAttack(&playerPos, &pos)
			t.checkAttackedByPlayer()
		}
	}

	t.handleAttackedState()

	if !t.attacked && t.player.Moving && !t.moving && distance > 70 {
		t.moving = true
		t.state = "running"
		t.UpdateState()
	}

	if !t.attacked && !strings.Contains(t.state, "attack") {
		if t.moving && t.state != "running" {
			t.state = "running"
			t.UpdateState()
		} else if !t.moving && t.state != "idle" {
			t.state = "idle"
			t.UpdateState()
		}
	}

	t.UpdateAnimation()
	println(t.state)
}

func (t *TorchGoblin) handleAimAttack(playerPos *rl.Vector2, pos *rl.Vector2) {
	var vec rl.Vector2 = rl.Vector2Subtract(*playerPos, *pos)
	var temp float32 = float32(math.Atan2(float64(vec.Y), float64(vec.X)) * 180 / math.Pi)
	var angle float32 = ternary(temp < 0, 360+temp, temp)

	if angle >= 315 || angle <= 45 {
		t.aimSide = ut.Globals.Side.Right
	} else if angle > 45 && angle < 135 {
		t.aimSide = ut.Globals.Side.Bottom
	} else if angle >= 135 && angle <= 225 {
		t.aimSide = ut.Globals.Side.Left
	} else if angle > 225 && angle < 315 {
		t.aimSide = ut.Globals.Side.Top
	}

	if !strings.Contains(t.state, "attack") {
		switch t.aimSide {
		case ut.Globals.Side.Left:
			t.state = "attack_left"
		case ut.Globals.Side.Right:
			t.state = "attack_right"
		case ut.Globals.Side.Bottom:
			t.state = "attack_bottom"
		case ut.Globals.Side.Top:
			t.state = "attack_top"
		}
		t.player.Attacked = true
		t.UpdateState()
	}
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
	anim.FrameCount = 6

	switch t.state {
	case "idle":
		anim.FrameCount = 7
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = 0
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	case "running":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	case "attack_left":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 2
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	case "attack_right":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 2
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	case "attack_bottom":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 3
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	case "attack_top":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 4
		t.sourceRec.X = float32(anim.FrameX)
		t.sourceRec.Y = float32(anim.FrameY)
	}
}

func (t *TorchGoblin) checkAttackedByPlayer() {
	if strings.Contains(t.player.state, "attack") && (t.player.animation.FrameX == 3 || t.player.animation.FrameX == 4) {
		if (t.player.aimSide == ut.Globals.Side.Top && t.aimSide == ut.Globals.Side.Bottom) ||
			(t.player.aimSide == ut.Globals.Side.Bottom && t.aimSide == ut.Globals.Side.Top) ||
			(t.player.aimSide == ut.Globals.Side.Left && t.aimSide == ut.Globals.Side.Right) ||
			(t.player.aimSide == ut.Globals.Side.Right && t.aimSide == ut.Globals.Side.Left) {
			t.attacked = true
			t.attackedForce = 1250
			t.attackedDirection = rl.Vector2Rotate(t.direction, math.Pi)
			t.state = "idle"
			t.UpdateState()
		}
	}
}

func (t *TorchGoblin) handleAttackedState() {
	if t.attacked {
		if t.attackedForce > 0 {
			t.Position = rl.Vector2Add(t.Position, rl.Vector2Scale(t.attackedDirection, t.attackedForce*rl.GetFrameTime()))
			t.destRec.X = t.Position.X
			t.destRec.Y = t.Position.Y
			t.CollisionBox.Rect.X = t.Position.X + 77
			t.CollisionBox.Rect.Y = t.Position.Y + 65
			t.attackedForce -= rl.GetFrameTime() * 8000
			if t.attackedForce < 0 {
				t.waitTimer = 0.120
			}
		} else {
			t.waitTimer -= rl.GetFrameTime()
			if t.waitTimer < 0 {
				t.attacked = false
			}
		}
	}
}

func (t *TorchGoblin) Draw() {
	if t.attacked {
		rl.BeginShaderMode(ut.Globals.Shaders.AttackedShader)
	}
	rl.DrawTexturePro(t.texture, t.sourceRec, t.destRec, t.origin, 0, rl.White)
	if t.attacked {
		rl.EndShaderMode()
	}
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
