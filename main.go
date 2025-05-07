package main

import (
	_ "fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 640
	BIRD_SIZE     = 60
)

var (
	LIGHT_BLUE = color.RGBA{173, 216, 230, 1}
)

type Coordinate struct {
	X, Y int
}

type Pipe struct {
	Coordinate
	Size int
}

type Game struct {
	BirdCoordinate Coordinate
	Pipes          []Pipe
	Score          int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(LIGHT_BLUE)
	ebitenutil.DebugPrint(screen, "oi")
	vector.DrawFilledRect(screen, float32(g.BirdCoordinate.X), float32(g.BirdCoordinate.Y), float32(BIRD_SIZE), float32(BIRD_SIZE), color.RGBA{250, 240, 0, 1}, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Flappy")
	game := Game{BirdCoordinate: Coordinate{X: WINDOW_WIDTH / 4, Y: WINDOW_HEIGHT / 2}, Pipes: []Pipe{}}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
