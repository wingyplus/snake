package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wingyplus/snake/mathutil"
)

type Snake struct {
	// (x, y) position default is (0, 0)
	x, y float64

	// (x, y) speed default is (1, 0)
	xs, ys float64

	scale float64
	imd   *imdraw.IMDraw
	win   *pixelgl.Window

	total int
	tail  []pixel.Vec
}

func (s *Snake) SetDirection(x, y float64) {
	s.xs = x
	s.ys = y
}

func (s *Snake) Update() {
	for i := 0; i < s.total-1; i++ {
		s.tail[i] = s.tail[i+1]
	}
	s.tail[s.total-1] = pixel.V(s.x, s.y)

	s.x = s.x + s.xs*s.scale
	s.y = s.y + s.ys*s.scale

	wx, wy := s.win.Bounds().Max.XY()

	s.x = mathutil.Constrain(s.x, 0, wx-s.scale)
	s.y = mathutil.Constrain(s.y, 0, wy-s.scale)
}

func (s *Snake) Show(win pixel.Target) {
	s.imd.Clear()
	for i := 0; i < s.total; i++ {
		t := s.tail[i]
		x, y := t.XY()

		s.imd.SetMatrix(pixel.IM.Moved(pixel.V(x, y)))
		s.imd.Push(pixel.V(0, 0), pixel.V(s.scale, s.scale))
		s.imd.Rectangle(0)
	}

	s.imd.Draw(win)
}

func (s *Snake) Pos() (x, y float64) {
	return s.x, s.y
}

func (s *Snake) Eat(food *Food) bool {
	fx, fy := food.Pos()
	sx, sy := s.Pos()
	d := mathutil.Dist(sx, sy, fx, fy)

	// when snake ate food, create tail by create new slice and copy from old slice to new slice
	if d < 1 {
		s.total++
		old := s.tail
		new := make([]pixel.Vec, s.total, s.total+1)
		copy(new, old)
		s.tail = new
		return true
	}
	return false
}

func NewSnake(scale float64, win *pixelgl.Window) *Snake {
	return &Snake{
		xs:    1,
		scale: scale,
		imd:   imdraw.New(nil),
		win:   win,
		total: 1,
		tail:  make([]pixel.Vec, 1),
	}
}
