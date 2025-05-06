package main

import (
	_ "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/rand/v2"
)

const (
	WIDTH      = 1280
	HEIGHT     = 640
	SNAKE_SIZE = 20
	FRUIT_SIZE = 20
	SPEED      = 1
)

type Direction int

const (
	Down Direction = iota
	Up
	Left
	Right
)

type Coordinate struct {
	X, Y int
}

type Game struct {
	SnakeBody       []Coordinate
	FruitCount      int
	FruitCoordinate Coordinate
	Direction
}

func (g *Game) Update() error {
	var x_speed, y_speed int
	switch g.Direction {
	case Left:
		x_speed, y_speed = -1, 0
	case Up:
		x_speed, y_speed = 0, -1
	case Down:
		x_speed, y_speed = 0, 1
	case Right:
		x_speed, y_speed = 1, 0
	}
	head := &g.SnakeBody[0]
	head.X += SPEED * x_speed * SNAKE_SIZE / 4
	head.Y += SPEED * y_speed * SNAKE_SIZE / 4

	// CheckCollision(g * Game)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for _, coord := range g.SnakeBody {
		vector.DrawFilledRect(screen, float32(coord.X), float32(coord.Y), float32(SNAKE_SIZE), float32(SNAKE_SIZE), color.White, false)
	}
	vector.DrawFilledRect(screen, float32(g.FruitCoordinate.X), float32(g.FruitCoordinate.Y), float32(FRUIT_SIZE), float32(FRUIT_SIZE), color.RGBA{255, 0, 0, 0}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Snake")
	start_x_pos, start_y_pos := rand.IntN(WIDTH), rand.IntN(HEIGHT)
	game := &Game{SnakeBody: []Coordinate{{X: 10, Y: 10}}, FruitCount: 0, FruitCoordinate: Coordinate{start_x_pos, start_y_pos}, Direction: Right}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
