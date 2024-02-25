package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/inconsolata"
)

const (
	screenWidth  = 640
	screenHeight = 320
	squareSize   = 64
	numSquares   = screenWidth / squareSize
)

var (
	squaresHeights        []float32
	squaresColor          []color.RGBA
	algorithms            []Algorithm
	currentAlgorithmIndex int
	sorted                = false

	bubbleSortN    = numSquares
	insertionSortI = 1
)

var lastUpdate time.Time = time.Now()

type Game struct{}

type Algorithm struct {
	name     string
	function func()
}

func GetRandomColor() color.RGBA {
	red := uint8(rand.Intn(256))
	green := uint8(rand.Intn(256))
	blue := uint8(rand.Intn(256))

	return color.RGBA{red, green, blue, 255}
}

func Swap(x, y int) {
	squaresHeights[x], squaresHeights[y] = squaresHeights[y], squaresHeights[x]
	squaresColor[x], squaresColor[y] = squaresColor[y], squaresColor[x]
}

func BubbleSortStep() {
	if bubbleSortN <= 1 || sorted {
		fmt.Println("Sorting done.")
		sorted = true
		return
	}
	newN := 0
	for i := 1; i < bubbleSortN; i++ {
		if squaresHeights[i-1] > squaresHeights[i] {
			Swap(i-1, i)
			newN = i
		}
	}
	bubbleSortN = newN

}

func InsertionSort() {
	i := 1
	for i < len(squaresHeights) {
		j := i
		for j > 0 && squaresHeights[j-1] > squaresHeights[j] {
			Swap(j, j-1)
			j--
		}
		i++
	}
}

func InsertionSortStep() {
	if insertionSortI >= len(squaresHeights) || sorted {
		fmt.Println("Sorting done.")
		sorted = true
		return
	}

	j := insertionSortI
	for j > 0 && squaresHeights[j-1] > squaresHeights[j] {
		Swap(j, j-1)
		j--
	}

	insertionSortI++
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && time.Since(lastUpdate) > time.Millisecond*350 {
		algorithms[currentAlgorithmIndex].function()
		lastUpdate = time.Now()
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && time.Since(lastUpdate) > time.Millisecond*350 {
		if currentAlgorithmIndex+1 >= len(algorithms) {
			currentAlgorithmIndex = 0
		} else {
			currentAlgorithmIndex++
		}

		lastUpdate = time.Now()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < numSquares; i++ {
		x := float32(i * squareSize)
		y := float32(screenHeight - squaresHeights[i]) //float32((screenHeight / 2) - squaresHeights[i])
		vector.DrawFilledRect(screen, x, y, squareSize, squaresHeights[i], squaresColor[i], false)
	}

	text.Draw(screen, algorithms[currentAlgorithmIndex].name, inconsolata.Regular8x16, 8, 16, color.White)
	// text.Draw(screen, "→: Next step\n↑: Next Algorithm", inconsolata.Regular8x16, 488, 16, color.White)
	text.Draw(screen, "_____________\n →: Next step\n↑: Next Algorithm", inconsolata.Regular8x16, 8, 24, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 320
}

func main() {
	squaresHeights = make([]float32, numSquares)
	squaresColor = make([]color.RGBA, numSquares)
	for i := 0; i < len(squaresHeights); i++ {
		squaresHeights[i] = rand.Float32() * (screenHeight - 1)
		squaresColor[i] = GetRandomColor()
	}

	algorithms = []Algorithm{{"Bubble Sort", BubbleSortStep}, {"Insertion Sort", InsertionSortStep}}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sorting Algorithm Visualizer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
