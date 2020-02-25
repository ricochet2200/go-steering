package main

import "github.com/fogleman/gg"

type Block struct {
	position Vector
	radius   float32
}

func MakeBlock(pos Vector) *Block {
	return &Block{pos, 20}
}

func (v *Block) Position() Vector {
	return v.position.Clone()
}
func (v *Block) Velocity() Vector {
	return Vector{0, 0, 0}
}
func (v *Block) MaxSpeed() float32 {
	return 0
}

func (v *Block) Radius() float32 {
	return v.radius
}

func (b *Block) Update() {

}
func (b *Block) Draw(c *gg.Context) {
	c.SetRGB(.5, .5, .5)
	c.DrawRectangle(float64(b.Position()[0]-b.Radius()), float64(b.Position()[1]-b.Radius()), float64(b.Position()[0]+b.Radius()), float64(b.Position()[1]+b.Radius()))
	c.Fill()
}
