package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	TargetPosition *rl.Vector2
	CameraPosition rl.Vector2
	Offset         rl.Vector2
}

func (c *Camera) Update() {
	c.Offset = rl.Vector2Subtract(c.CameraPosition, *c.TargetPosition)
	c.TargetPosition.X = c.CameraPosition.X
	c.TargetPosition.Y = c.CameraPosition.Y
}
