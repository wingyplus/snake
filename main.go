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

const (
	scale  = 20
	width  = 600
	height = 600
)

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "The Snake Game",
		Bounds: pixel.R(0, 0, width, height),
	})
	if err != nil {
		log.Fatal(err)
	}

	x, y := pickLocation()
	food := NewFood(scale)
	food.SetLocation(x, y)
	snake := NewSnake(scale, win)

	for !win.Closed() {
		handleInput(win, snake)
		win.Clear(color.Gray{51})
		snake.Update()
		snake.Show(win)
		food.Update()
		food.Show(win)

		if snake.Eat(food) {
			x, y = pickLocation()
			food.SetLocation(x, y)
		}

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

func pickLocation() (x, y float64) {
	vec := pixel.V(float64(rand.Intn(width/scale)), float64(rand.Intn(height/scale)))
	vec = vec.ScaledXY(pixel.V(scale, scale))
	x = vec.X()
	y = vec.Y()
	return
}
