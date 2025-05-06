package main

import (
	_ "fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"strconv"

	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (g *Game) OutOfBounds() bool {
	head := &g.SnakeBody[0]
	return head.X < 0 || head.X > WIDTH-SNAKE_SIZE || head.Y < 0 || head.Y >= HEIGHT-SNAKE_SIZE
}

func (g *Game) CheckCollisionWithFruit() bool {
	head := &g.SnakeBody[0]
	a, b := head.X-g.FruitCoordinate.X, head.Y-g.FruitCoordinate.Y
	if a*a+b*b <= SNAKE_SIZE*SNAKE_SIZE {
		return true
	}
	return false
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

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		g.Direction = Up
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		g.Direction = Down
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		g.Direction = Left
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		g.Direction = Right
	}
	if g.OutOfBounds() {
		return errors.New("Snake died colliding with wall!")
	}
	if g.CheckCollisionWithFruit() {
		g.FruitCount += 1
		x_pos, y_pos := rand.IntN(WIDTH/FRUIT_SIZE)*FRUIT_SIZE, rand.IntN(HEIGHT/FRUIT_SIZE)*FRUIT_SIZE
		g.FruitCoordinate = Coordinate{X: x_pos, Y: y_pos}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(g.FruitCount), 0, 0)
	// draw snake
	for _, coord := range g.SnakeBody {
		vector.DrawFilledRect(screen, float32(coord.X), float32(coord.Y), float32(SNAKE_SIZE), float32(SNAKE_SIZE), color.RGBA{0, 255, 0, 1}, false)
	}
	// draw fruit
	vector.DrawFilledRect(screen, float32(g.FruitCoordinate.X), float32(g.FruitCoordinate.Y), float32(FRUIT_SIZE), float32(FRUIT_SIZE), color.RGBA{255, 0, 0, 0}, false)
	// draw grid
	stoke_width := float32(1)
	for i := range WIDTH / SNAKE_SIZE {
		vector.StrokeLine(screen, float32(i*SNAKE_SIZE), float32(0), float32(i*SNAKE_SIZE), float32(HEIGHT), stoke_width, color.RGBA{173, 173, 173, 0}, false)
	}
	for i := range HEIGHT / SNAKE_SIZE {
		vector.StrokeLine(screen, float32(0), float32(i*SNAKE_SIZE), float32(WIDTH), float32(i*SNAKE_SIZE), stoke_width, color.RGBA{173, 173, 173, 0}, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Snake")
	start_x_pos, start_y_pos := rand.IntN(WIDTH/FRUIT_SIZE)*FRUIT_SIZE, rand.IntN(HEIGHT/FRUIT_SIZE)*FRUIT_SIZE
	game := Game{SnakeBody: []Coordinate{{X: 20, Y: 20}}, FruitCount: 0, FruitCoordinate: Coordinate{start_x_pos, start_y_pos}, Direction: Right}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
