package sprites

import (
	ut "main/utils"
	"math"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Knight struct {
	Position       rl.Vector2
	direction      rl.Vector2
	origin         rl.Vector2
	texture        rl.Texture2D
	animation      ut.Animation
	sourceRec      rl.Rectangle
	destRec        rl.Rectangle
	state          string
	Speed          int
	attack_no      int
	attack_counter int
	mouse_timer    float32
	CollisionBox   ut.CollisionBox
	aimSide        int
}

func (k *Knight) Load(texture rl.Texture2D) {
	k.state = "idle"
	k.texture = texture
	k.Speed = 400
	k.direction = rl.Vector2Zero()
	k.animation = ut.Animation{
		FrameX:      0,
		FrameCount:  6,
		FrameWidth:  uint16(k.texture.Width) / 6,
		FrameHeight: uint16(k.texture.Height) / 8,
		FrameY:      0,
		FrameTime:   0.100,
	}
	k.Position = rl.NewVector2(float32(ut.Globals.CanvasWidth)/2-float32(k.animation.FrameWidth)/2, float32(ut.Globals.CanvasHeight)/2-float32(k.animation.FrameHeight)/2)
	k.sourceRec = rl.Rectangle{
		X:      float32(k.animation.FrameX),
		Y:      float32(k.animation.FrameY),
		Width:  float32(k.animation.FrameWidth),
		Height: float32(k.animation.FrameHeight),
	}
	k.destRec = rl.Rectangle{
		X:      float32(k.animation.FrameX),
		Y:      float32(k.animation.FrameY),
		Width:  float32(k.animation.FrameWidth),
		Height: float32(k.animation.FrameHeight),
	}
	k.destRec.X = k.Position.X
	k.destRec.Y = k.Position.Y
	k.origin = rl.Vector2{}
	k.attack_no = 0
	k.attack_counter = 0
	k.mouse_timer = 0.500
	k.CollisionBox = ut.CollisionBox{Rect: rl.NewRectangle(k.Position.X+74, k.Position.Y+68, 41, 60)}
}

func (k *Knight) UnLoad() {

}

func (k *Knight) Update() {
	k.handleMouse()
	k.handleAttack()
	if !strings.Contains(k.state, "attack") {
		if rl.IsKeyDown(rl.KeyD) {
			k.direction.X = 1
			if k.sourceRec.Width < 0 {
				k.sourceRec.Width *= -1
			}
		} else if rl.IsKeyDown(rl.KeyA) {
			k.direction.X = -1
			if k.sourceRec.Width > 0 {
				k.sourceRec.Width *= -1
			}
		} else {
			k.direction.X = 0
		}

		if rl.IsKeyDown(rl.KeyS) {
			k.direction.Y = 1
		} else if rl.IsKeyDown(rl.KeyW) {
			k.direction.Y = -1
		} else {
			k.direction.Y = 0
		}

		k.direction = rl.Vector2Normalize(k.direction)

		if ((k.direction.X != 0) || (k.direction.Y != 0)) && (k.state != "run") {
			k.state = "run"
			k.UpdateState()
		} else if (k.direction.X == 0) && (k.direction.Y == 0) && (k.state != "idle") {
			k.state = "idle"
			k.UpdateState()
		}

		if (k.direction.X != 0) || (k.direction.Y != 0) {
			k.Position = rl.Vector2Add(k.Position, rl.Vector2Scale(k.direction, float32(k.Speed)*rl.GetFrameTime()))
			k.destRec.X = k.Position.X
			k.destRec.Y = k.Position.Y
			k.CollisionBox.Rect.X = k.Position.X + 74
			k.CollisionBox.Rect.Y = k.Position.Y + 68
		}
	}

	k.UpdateAnimation()
}

func (k *Knight) Draw() {
	rl.DrawTexturePro(k.texture, k.sourceRec, k.destRec, k.origin, 0, rl.White)
	// rl.DrawRectangleLinesEx(k.CollisionBox.Rect, 1, rl.Red)
}

func (k *Knight) UpdateAnimation() {
	var anim *ut.Animation = &k.animation
	anim.FrameTime -= rl.GetFrameTime()
	if anim.FrameTime < 0.0 {
		anim.FrameTime = 0.100
		anim.FrameX++
		if anim.FrameX == uint16(anim.FrameCount) {
			anim.FrameX = 0
			if strings.Contains(k.state, "attack") {
				k.state = "idle"
				k.UpdateState()
			}
		}
		k.sourceRec.X = float32(anim.FrameX) * float32(anim.FrameWidth)
	}
}

func (k *Knight) UpdateState() {
	var anim *ut.Animation = &k.animation
	switch k.state {
	case "idle":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = 0
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "run":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "side_attack1":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 2
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "side_attack2":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 3
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "down_attack1":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 4
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "down_attack2":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 5
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "up_attack1":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 6
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	case "up_attack2":
		anim.FrameTime = 0.100
		anim.FrameX = 0
		anim.FrameY = anim.FrameHeight * 7
		k.sourceRec.X = float32(anim.FrameX)
		k.sourceRec.Y = float32(anim.FrameY)
	}
}

func (k *Knight) handleMouse() {
	var vec rl.Vector2 = rl.Vector2Subtract(ut.Globals.MousePos, rl.NewVector2(k.Position.X+float32(k.animation.FrameWidth)/2, k.Position.Y+float32(k.animation.FrameHeight)/2))
	var temp float32 = float32(math.Atan2(float64(vec.Y), float64(vec.X)) * 180 / math.Pi)
	var angle float32 = ternary(temp < 0, 360+temp, temp)

	if angle >= 315 || angle <= 45 {
		k.aimSide = ut.Globals.Side.Right
	} else if angle > 45 && angle < 135 {
		k.aimSide = ut.Globals.Side.Bottom
	} else if angle >= 135 && angle <= 225 {
		k.aimSide = ut.Globals.Side.Left
	} else if angle > 225 && angle < 315 {
		k.aimSide = ut.Globals.Side.Top
	}
}

func (k *Knight) handleAttack() {
	if k.mouse_timer > 0 {
		k.mouse_timer -= rl.GetFrameTime()
	}

	if k.mouse_timer < 0 {
		k.attack_counter = 0
		k.attack_no = 0
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		k.mouse_timer = 0.500
		k.attack_counter++
	}

	if k.attack_counter != 0 && !strings.Contains(k.state, "attack") {
		rl.PlaySound(ut.Globals.Sound.PlayerAttackSound)
		k.attack_counter--
		if k.attack_no == 2 {
			k.attack_no = 0
		}
		if k.attack_no == 0 {
			switch k.aimSide {
			case ut.Globals.Side.Right:
				if k.sourceRec.Width < 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "side_attack1"
			case ut.Globals.Side.Left:
				if k.sourceRec.Width > 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "side_attack1"
			case ut.Globals.Side.Top:
				if k.sourceRec.Width < 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "up_attack1"
			case ut.Globals.Side.Bottom:
				if k.sourceRec.Width < 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "down_attack1"
			}
			k.UpdateState()
		} else {
			switch k.aimSide {
			case ut.Globals.Side.Right:
				if k.sourceRec.Width < 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "side_attack2"
			case ut.Globals.Side.Left:
				if k.sourceRec.Width > 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "side_attack2"
			case ut.Globals.Side.Top:
				if k.sourceRec.Width < 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "up_attack2"
			case ut.Globals.Side.Bottom:
				if k.sourceRec.Width < 0 {
					k.sourceRec.Width *= -1
				}
				k.state = "down_attack2"
			}
			k.UpdateState()
		}
		k.attack_no++
	}
}
