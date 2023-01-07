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
	//go:embed resources/projectile.png
	projectilePng    []byte
	projectileSprite *ebiten.Image
)

type Projectiles struct {
	instances []*projInstance
}

type projInstance struct {
	x        float64
	y        float64
	dir      f64.Vec2
	rotation int
}

func init() {
	// Get image here
	img, err := png.Decode(bytes.NewReader(projectilePng))
	if err != nil {
		log.Fatal(err)
	}

	projectileSprite = ebiten.NewImageFromImage(img)
}

func (p *Projectiles) New(x, y float64, dir f64.Vec2, rot int) {
	p.instances = append(p.instances, &projInstance{
		x:        x,
		y:        y,
		dir:      dir,
		rotation: rot,
	})
}

func (p *Projectiles) Update(g *Game) error {
	for _, proj := range g.projectiles.instances {
		mag := GetMag(proj.dir)

		proj.x -= proj.dir[0] / mag * 5
		proj.y -= proj.dir[1] / mag * 5

		// fmt.Printf("%d: %f:%f\n", i, proj.x, proj.y)
	}

	return nil
}

func (p *Projectiles) Draw(screen *ebiten.Image) {
	for _, proj := range p.instances {
		w, h := projectileSprite.Size()
		op := &ebiten.DrawImageOptions{}

		// Move the image's center to the screen's upper-left corner.
		// This is a preparation for rotating. When geometry matrices are applied,
		// the origin point is the upper-left corner.
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		// Rotate the image. As a result, the anchor point of this rotate is
		// the center of the image.
		op.GeoM.Rotate(float64(proj.rotation%360) * 2 * math.Pi / 360)

		// Move the image to the screen's center.
		op.GeoM.Translate(proj.x, proj.y)
		op.GeoM.Scale(1, 1)

		screen.DrawImage(projectileSprite, op)
	}
}

func GetMag(vec2 f64.Vec2) float64 {
	return math.Sqrt(vec2[0]*vec2[0] + vec2[1]*vec2[1])
}
