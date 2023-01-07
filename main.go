package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LD52")
	if err := ebiten.RunGame(&Game{projectiles: Projectiles{instances: []*projInstance{}}}); err != nil {
		log.Fatal(err)
	}
}
