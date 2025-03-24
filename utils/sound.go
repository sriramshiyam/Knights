package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Sounds struct {
	PlayerAttackSound rl.Sound
}

func (s *Sounds) Load() {
	s.PlayerAttackSound = rl.LoadSound("res/sound/sword.wav")
}

func (s *Sounds) UnLoad() {
	rl.UnloadSound(s.PlayerAttackSound)
}
