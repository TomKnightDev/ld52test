package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 768
)

type Game struct {
	turret      Turret
	projectiles Projectiles
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.turret.Update(g)
	g.projectiles.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.turret.Draw(screen)
	g.projectiles.Draw(screen)
}
