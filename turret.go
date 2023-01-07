package main

import (
	"bytes"
	_ "embed"
	"image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

var (
	//go:embed resources/turret.png
	turretPng                     []byte
	turretSprite                  *ebiten.Image
	actionTime, currentActionTime int
)

func init() {
	// Get image here
	img, err := png.Decode(bytes.NewReader(turretPng))
	if err != nil {
		log.Fatal(err)
	}

	turretSprite = ebiten.NewImageFromImage(img)
	actionTime = 10
}

type Turret struct {
	pos      f64.Vec2
	rotation int
}

func (t *Turret) Update(g *Game) error {
	t.pos = f64.Vec2{ScreenWidth / 2, ScreenHeight / 2}
	x, y := ebiten.CursorPosition()

	xf := float64(x) - t.pos[0]
	yf := float64(y) - t.pos[1]

	angle := math.Atan2(yf, xf) * (180 / math.Pi)

	t.rotation = int(angle) + 90

	currentActionTime++
	if currentActionTime >= actionTime && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		currentActionTime = 0
		xdir := t.pos[0] - float64(x)
		ydir := t.pos[1] - float64(y)

		g.projectiles.New(ScreenWidth/2, ScreenHeight/2, f64.Vec2{xdir, ydir}, t.rotation)
	}

	return nil
}

func (t *Turret) Draw(screen *ebiten.Image) {
	w, h := turretSprite.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.
	op.GeoM.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)

	// Move the image to the screen's center.
	op.GeoM.Translate(t.pos[0], t.pos[1])
	op.GeoM.Scale(1, 1)

	screen.DrawImage(turretSprite, op)

	// screen.DrawImage(turretSprite, GetDrawOptions(100, 100, t.rotation))
}

func GetDrawOptions(transX, transY float64, rotation int) *ebiten.DrawImageOptions {
	m := ebiten.GeoM{}
	m.Scale(1, 1)
	m.Translate(transX, transY)
	m.Rotate(float64(rotation%360) * 2 * math.Pi / 360)

	return &ebiten.DrawImageOptions{
		GeoM: m,
	}
}
