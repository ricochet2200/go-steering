package steering

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/fogleman/gg"
)

type Prey struct {
	*Vehicle
	Sphere        float32
	steeringForce Vector
	count         int
	level         *Level
	perception    float32
}

func MakePrey(l *Level) *Prey {
	v := MakeVehicle(Vector{rand.Float32() * 200, rand.Float32() * 200, 0}, 10, 10, 5, 10)
	return &Prey{v, 5, Vector{0, 0, 0}, 0, l, 100}
}

func (p *Prey) Update() {
	fmt.Println("Prey Update")
	p.count -= 1
	//p.count = 2

	if contain := Contain(p, Vector{0, 0, 0}, p.level.Bounds(), p.perception, 2); contain != nil {
		fmt.Println("Contain")
		p.steeringForce = contain
	} else if avoid := Avoid(p, p.level.Vehicles, p.perception); avoid != nil {
		fmt.Println("Avoid")
		p.steeringForce = avoid
	} else {
		if p.count <= 0 {
			log.Println("Wander")
			p.steeringForce = Wander(p, p.Sphere, p.Orientation)
			p.count = 100
		}
	}
	p.Move(p.steeringForce)
}

func (p *Prey) Draw(c *gg.Context) {
	// Prey
	c.SetRGB(.5, .5, .5)
	c.DrawCircle(float64(p.X()), float64(p.Y()), float64(p.Radius()))
	c.Fill()

	c.SetRGB(0, 0, 1)
	c.DrawLine(float64(p.X()), float64(p.Y()), float64(p.X()+p.steeringForce[0]*p.MaxForce), float64(p.Y()+p.steeringForce[1]*p.MaxForce))
	c.Stroke()

	// Forward perception
	vision := p.Forward().Mult(p.perception)
	c.SetRGB(1, 0, 1)
	c.DrawLine(float64(p.X()), float64(p.Y()), float64(p.X()+vision[0]), float64(p.Y()+vision[1]))
	c.Stroke()

	closest := Vector(nil)
	closestDistance := float32(100000000)

	for _, sphere := range p.level.Vehicles {
		rel := sphere.Position().Minus(p.Position())
		distance := rel.Len()
		if distance < closestDistance && distance > 0 && distance < p.perception {
			closestDistance = distance
			closest = sphere.Position()

		}

	}

	if closest != nil {
		c.SetRGB(.1, .1, .1)
		c.DrawLine(float64(p.X()), float64(p.Y()), float64(closest[0]), float64(closest[1]))
		c.Stroke()
	}

}
