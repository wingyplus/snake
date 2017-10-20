package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pixelgl.Run(run)
}

const scale = 20

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "The Snake Game",
		Bounds: pixel.R(0, 0, 600, 600),
	})
	if err != nil {
		log.Fatal(err)
	}

	x, y := randPositionXY()
	food := NewFood(x, y, scale)
	snake := NewSnake(scale, win)

	for !win.Closed() {
		handleInput(win, snake)
		win.Clear(color.Gray{51})
		snake.Update()
		snake.Show(win)
		food.Update()
		food.Show(win)
		win.Update()
	}
}

func handleInput(win *pixelgl.Window, snake *Snake) {
	if win.JustPressed(pixelgl.KeyLeft) {
		snake.SetDirection(-1, 0)
	} else if win.JustPressed(pixelgl.KeyRight) {
		snake.SetDirection(1, 0)
	} else if win.JustPressed(pixelgl.KeyUp) {
		snake.SetDirection(0, 1)
	} else if win.JustPressed(pixelgl.KeyDown) {
		snake.SetDirection(0, -1)
	}
}

func randPositionXY() (x, y float64) {
	x = float64(rand.Intn(600))
	y = float64(rand.Intn(600))
	return
}
