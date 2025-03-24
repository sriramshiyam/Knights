package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Side struct {
	Right  int
	Bottom int
	Left   int
	Top    int
}

var Globals = struct {
	CanvasWidth  int
	CanvasHeight int
	CanavsDest   rl.Rectangle
	CanvasSource rl.Rectangle
	MousePos     rl.Vector2
	Sound        Sounds
	Side         Side
}{1152, 648, rl.Rectangle{}, rl.Rectangle{}, rl.Vector2{}, Sounds{}, Side{0, 1, 2, 3}}
