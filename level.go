package steering

import (
	"fmt"

	"github.com/fogleman/gg"
)

type Level struct {
	Width    int
	Height   int
	Vehicles []Entity
	spawn    Spawn
}

func MakeLevel(width int, height int) *Level {

	fmt.Println(width, height)
	e := []Entity{}
	level := &Level{width, height, e, Spawn{}}
	level.spawn.level = level
	level.spawn.position = Vector{50, 50, 0}

	/*	for i := 0; i < int(width/40); i++ {
			e = append(e, MakeBlock(Vector{float32(i*40 + 20), 20, 0}))                   // TOP
			e = append(e, MakeBlock(Vector{float32(i*40 + 20), float32(height - 20), 0})) // Bottom
		}
		for i := 0; i < int(height/40); i++ {
			e = append(e, MakeBlock(Vector{20, float32(i*40 + 20), 0})) // LEFT
			e = append(e, MakeBlock(Vector{float32(width - 20), float32(i*40 + 20), 0}))
		}*/

	p := MakePrey(level)
	e = append(e, p)
	//	e = append(e, MakeHunter(p))
	for i := 0; i < 20; i++ {
		e = append(e, MakePrey(level))
	}
	fmt.Println("Number of Objects: ", len(e))
	level.Vehicles = e
	return level
}

func (l *Level) Bounds() Vector {
	return Vector{float32(l.Width), float32(l.Height), 0}
}

func (l *Level) Update() {
	l.spawn.Update()
	for i, _ := range l.Vehicles {
		l.Vehicles[i].Update()
	}
}

func (l *Level) Draw(c *gg.Context) {
	for i, _ := range l.Vehicles {
		l.Vehicles[i].Draw(c)
	}
}
