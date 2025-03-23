package utils

import rl "github.com/gen2brain/raylib-go/raylib"

var Globals = struct {
	CanvasWidth  int
	CanvasHeight int
	CanavsDest   rl.Rectangle
	CanvasSource rl.Rectangle
	MousePos     rl.Vector2
}{1152, 648, rl.Rectangle{}, rl.Rectangle{}, rl.Vector2{}}
