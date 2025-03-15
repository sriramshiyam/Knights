package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type AABBInfo struct {
	OverlapX         float32
	OverlapY         float32
	PreviousOverlapX float32
	PreviousOverlapY float32
}

type CollisionBox struct {
	Rect     rl.Rectangle
	AABBInfo AABBInfo
}
