package steering

import "github.com/fogleman/gg"

type Entity interface {
	Update()
	Draw(*gg.Context)
	Position() Vector
	Radius() float32
}
