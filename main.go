package main

import (
	"image/color"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "The Snake Game",
		Bounds: pixel.R(0, 0, 600, 600),
	})
	if err != nil {
		log.Fatal(err)
	}

	snake := NewSnake(win)

	for !win.Closed() {
		handleInput(win, snake)
		win.Clear(color.Gray{51})
		snake.Update()
		snake.Show(win)
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
