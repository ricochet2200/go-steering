package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

type Hunter struct {
	*Vehicle
	quarry Quarry
}

func MakeHunter(q Quarry) *Hunter {
	v := MakeVehicle(Vector{10, 10, 0}, 20, 5, 2, 10)
	return &Hunter{v, q}
}

func (h *Hunter) Update() {
	fmt.Println("hunter update")
	h.Move(Pursuit(h, h.quarry, 10))
}

func (h *Hunter) Draw(c *gg.Context) {
	c.SetRGB(0, 0, 0)
	c.DrawCircle(float64(h.X()), float64(h.Y()), float64(h.Radius()/2))
	c.Fill()

	c.DrawLine(float64(h.X()), float64(h.Y()), float64(h.X()+h.Radius()*h.Orientation[0][0]), float64(h.Y()+h.Radius()*h.Orientation[0][1]))
	c.SetRGB(1, 0, 0)
	c.Stroke()
}
