package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Shaders struct {
	AttackedShader rl.Shader
}

func (s *Shaders) Load() {
	s.AttackedShader = rl.LoadShader("", "res/shader/attackedshader.fx")
}

func (s *Shaders) UnLoad() {
	rl.UnloadShader(s.AttackedShader)
}
