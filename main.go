package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 640
	squareSize   = 64
	numSquares   = screenWidth / squareSize
)

var (
	squaresHeights []float32
	squaresColor   []color.RGBA
	bubbleSortN    = numSquares
)

var lastUpdate time.Time = time.Now()

type Game struct{}

func getRandomColor() color.RGBA {
	red := uint8(rand.Intn(256))
	green := uint8(rand.Intn(256))
	blue := uint8(rand.Intn(256))

	return color.RGBA{red, green, blue, 255}
}

func swap(x, y int) {
	squaresHeights[x], squaresHeights[y] = squaresHeights[y], squaresHeights[x]
	squaresColor[x], squaresColor[y] = squaresColor[y], squaresColor[x]
}

func bubbleSortStep() {
	newN := 0
	for i := 1; i < bubbleSortN; i++ {
		if squaresHeights[i-1] > squaresHeights[i] {
			swap(i-1, i)
			newN = i
		}
	}
	bubbleSortN = newN

}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && time.Since(lastUpdate) > time.Millisecond*350 {
		if bubbleSortN > 1 {
			bubbleSortStep()
		} else {
			fmt.Println("Sorting done.")
		}
		lastUpdate = time.Now()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < numSquares; i++ {
		x := float32(i * squareSize)
		y := float32((screenHeight / 2) - squaresHeights[i])
		vector.DrawFilledRect(screen, x, y, squareSize, squaresHeights[i], squaresColor[i], false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	squaresHeights = make([]float32, numSquares)
	squaresColor = make([]color.RGBA, numSquares)
	for i := 0; i < len(squaresHeights); i++ {
		squaresHeights[i] = rand.Float32() * ((screenHeight / 2) - 1)
		squaresColor[i] = getRandomColor()
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sorting Algorithm Visualizer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
