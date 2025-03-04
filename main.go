package main

import (
	sp "main/sprites"
	ut "main/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Knights")
	rl.InitAudioDevice()
	rl.MaximizeWindow()
	rl.SetTargetFPS(60)

	ut.Globals.CanavsDest.Width = float32(ut.Globals.CanvasWidth)
	ut.Globals.CanavsDest.Height = float32(ut.Globals.CanvasHeight)
	ut.Globals.CanvasSource.Width = float32(ut.Globals.CanvasWidth)
	ut.Globals.CanvasSource.Height = -float32(ut.Globals.CanvasHeight)

	var Canvas rl.RenderTexture2D = rl.LoadRenderTexture(int32(ut.Globals.CanvasWidth), int32(ut.Globals.CanvasHeight))

	var music rl.Music = rl.LoadMusicStream("res/music/country.mp3")
	music.Looping = true
	rl.PlayMusicStream(music)

	var knight sp.Knight = sp.Knight{}
	knight.Load()
	var ground sp.Ground = sp.Ground{}
	ground.Load()
	var camera ut.Camera = ut.Camera{}
	camera.TargetPosition = &knight.Position
	camera.CameraPosition = knight.Position

	var clearColor rl.Color = rl.NewColor(90, 169, 167, 255)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)
		if rl.IsWindowResized() {
			ResizeCanvas()
		}
		knight.Update()
		camera.Update()
		ground.ApplyCameraOffset(camera.Offset)

		rl.BeginTextureMode(Canvas)

		rl.ClearBackground(clearColor)

		// rl.DrawFPS(0, 0)
		ground.Draw()
		knight.Draw()

		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexturePro(Canvas.Texture, ut.Globals.CanvasSource, ut.Globals.CanavsDest, rl.Vector2{}, 0, rl.White)
		rl.EndDrawing()
	}

	knight.UnLoad()
	ground.UnLoad()
	rl.UnloadRenderTexture(Canvas)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func ResizeCanvas() {
	var screenWidth int = rl.GetScreenWidth()
	var screenHeight int = rl.GetScreenHeight()

	var targetAspect float32 = float32(ut.Globals.CanvasWidth) / float32(ut.Globals.CanvasHeight)
	var windowAspect float32 = float32(screenWidth) / float32(screenHeight)

	var canvasDest *rl.Rectangle = &ut.Globals.CanavsDest
	if targetAspect > windowAspect {
		canvasDest.Width = float32(screenWidth)
		canvasDest.Height = float32(screenWidth) / targetAspect
		canvasDest.X = 0
		canvasDest.Y = (float32(screenHeight) - canvasDest.Height) / 2
	} else {
		canvasDest.Width = float32(screenHeight) * targetAspect
		canvasDest.Height = float32(screenHeight)
		canvasDest.X = (float32(screenWidth) - canvasDest.Width) / 2
		canvasDest.Y = 0
	}
}
