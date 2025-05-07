package main

import (
	"errors"
	_ "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/rand/v2"
	"strconv"
	"time"
)

const (
	WINDOW_WIDTH            = 1280
	WINDOW_HEIGHT           = 640
	BIRD_SIZE               = 60
	INITIAL_BIRD_JUMP_SPEED = 10
	GRAVITY                 = 1
	PIPE_WIDTH              = 100
	PIPE_EMPTY_SIZE         = 200
	PIPE_SPEED              = 5
)

var (
	LIGHT_BLUE   = color.RGBA{173, 216, 230, 255}
	GREEN        = color.RGBA{0, 255, 0, 255}
	AMBER_YELLOW = color.RGBA{255, 255, 0, 255}
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
	BirdSpeed      int
	Pipes          []Pipe
	Score          int
	LastAddedPipe  time.Time
}

func (g *Game) OutOfBounds() bool {
	return g.BirdCoordinate.X < 0 || g.BirdCoordinate.X > WINDOW_WIDTH || g.BirdCoordinate.Y < 0 || g.BirdCoordinate.Y > WINDOW_HEIGHT
}

func (g *Game) CheckCollision() bool {
	return true
}

func (g *Game) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeySpace):
		g.BirdSpeed = INITIAL_BIRD_JUMP_SPEED
	default:
	}
	g.BirdCoordinate.Y -= g.BirdSpeed
	g.BirdSpeed -= GRAVITY
	if g.OutOfBounds() {
		return errors.New("Bird out of bounds!")
	}
	for i, _ := range g.Pipes {
		g.Pipes[i].X -= PIPE_SPEED
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(LIGHT_BLUE)
	ebitenutil.DebugPrint(screen, strconv.Itoa(g.Score))
	// draw bird
	vector.DrawFilledRect(screen, float32(g.BirdCoordinate.X), float32(g.BirdCoordinate.Y), float32(BIRD_SIZE), float32(BIRD_SIZE), AMBER_YELLOW, false)
	// draw pipes
	for _, pipe := range g.Pipes {
		x, y := pipe.X, pipe.Y
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(PIPE_WIDTH), float32(pipe.Size), GREEN, false)
		vector.DrawFilledRect(screen, float32(x), float32(y+pipe.Size+PIPE_EMPTY_SIZE), float32(PIPE_WIDTH), float32(WINDOW_HEIGHT-pipe.Size-PIPE_EMPTY_SIZE), GREEN, false)
	}
	upper_pipe_size := rand.IntN(WINDOW_HEIGHT / 2)
	if time.Since(g.LastAddedPipe) > time.Second*2 {
		g.Pipes = append(g.Pipes, Pipe{Coordinate{X: WINDOW_WIDTH, Y: 0}, upper_pipe_size})
		g.LastAddedPipe = time.Now()
	}

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
