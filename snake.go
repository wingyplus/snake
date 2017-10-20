package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wingyplus/codingchallenge/snake/mathutil"
)

type Snake struct {
	// (x, y) position default is (0, 0)
	x, y float64

	// (x, y) speed default is (1, 0)
	xs, ys float64

	scale float64
	imd   *imdraw.IMDraw
	win   *pixelgl.Window
}

func (s *Snake) SetDirection(x, y float64) {
	s.xs = x
	s.ys = y
}

func (s *Snake) Update() {
	s.x = s.x + s.xs*s.scale*0.01
	s.y = s.y + s.ys*s.scale*0.01

	wx, wy := s.win.Bounds().Max.XY()

	s.x = mathutil.Constrain(s.x, 0, wx-s.scale)
	s.y = mathutil.Constrain(s.y, 0, wy-s.scale)
}

func (s *Snake) Show(win pixel.Target) {
	s.imd.Clear()
	s.imd.SetMatrix(pixel.IM.Moved(pixel.V(s.x, s.y)))
	s.imd.Push(pixel.V(0, 0), pixel.V(s.scale, s.scale))
	s.imd.Rectangle(0)
	s.imd.Draw(win)
}

func NewSnake(win *pixelgl.Window) *Snake {
	return &Snake{
		xs:    1,
		scale: 20,
		imd:   imdraw.New(nil),
		win:   win,
	}
}