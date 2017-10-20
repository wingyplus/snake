package main

import (
	"image/color"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Snake struct {
	// (x, y) position default is (0, 0)
	x, y float64

	// (x, y) speed default is (1, 0)
	xs, ys float64

	imd *imdraw.IMDraw
}

func (s *Snake) Update() {
	s.x = s.x + s.xs
	s.y = s.y + s.ys
}

func (s *Snake) Show(win pixel.Target) {
	s.imd.Clear()
	s.imd.SetMatrix(pixel.IM.Moved(pixel.V(s.x, s.y)))
	s.imd.Push(pixel.V(0, 0), pixel.V(10, 10))
	s.imd.Rectangle(0)
	s.imd.Draw(win)
}

func NewSnake() *Snake {
	return &Snake{
		xs:  1,
		imd: imdraw.New(nil),
	}
}

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

	snake := NewSnake()

	for !win.Closed() {
		win.Clear(color.Gray{51})
		snake.Update()
		snake.Show(win)
		win.Update()
	}
}
