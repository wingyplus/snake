package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Food struct {
	x, y  float64
	scale float64
	imd   *imdraw.IMDraw
}

func (f *Food) Update() {}

func (f *Food) Show(win pixel.Target) {
	f.imd.Clear()
	f.imd.Color(colornames.Red)
	f.imd.SetMatrix(pixel.IM.Moved(pixel.V(f.x, f.y)))
	f.imd.Push(pixel.V(0, 0), pixel.V(f.scale, f.scale))
	f.imd.Rectangle(0)
	f.imd.Draw(win)
}

func (f *Food) SetLocation(x, y float64) {
	f.x = x
	f.y = y
}

func NewFood(scale float64) *Food {
	return &Food{
		scale: scale,
		imd:   imdraw.New(nil),
	}
}
