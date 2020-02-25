package steering

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/gg"
)

type Spawn struct {
	level    *Level
	position Vector
}

func (s *Spawn) Update() {
	if rand.Intn(50) == -1 {
		fmt.Println("adding a prey")
		s.level.Vehicles = append(s.level.Vehicles, MakePrey(s.level))
	}
}

func (s *Spawn) Draw(c *gg.Context) {
	for i, _ := range s.level.Vehicles {
		s.level.Vehicles[i].Draw(c)
	}
}
