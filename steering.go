package main

import (
	"fmt"
	"math/rand"
)

func Seek(s Steerer, target Vector) Vector {
	desired_velocity := target.Minus(s.Position()).Normalize().Mult(s.MaxSpeed())
	return desired_velocity.Minus(s.Velocity())
}

func Flee(s Steerer, target Vector) Vector {
	desired_velocity := target.Minus(s.Position()).Normalize().Mult(s.MaxSpeed())
	return s.Velocity().Minus(desired_velocity)
}

func Pursuit(s Steerer, q Quarry, futureUpdates float32) Vector {
	pred := q.Position().Plus(q.Velocity().Mult(futureUpdates))
	return Seek(s, pred)
}

func Wander(s Steerer, sphere float32, orientation Matrix) Vector {
	x := rand.Float32()
	y := rand.Float32()
	z := rand.Float32()
	v := Vector{x, y, z}
	//	fmt.Println("wander", v.Normalize(), sphere, orientation)

	force := v.Normalize().Mult(sphere)
	//	fmt.Println("undirected force", force)
	return orientation.Mult(force)
}

func Avoid(s Steerer, spheres []Entity, length float32) Vector {
	closest := Entity(nil)
	closestDistance := float32(100000000)
	pos := s.Position()
	for _, sphere := range spheres {
		rel := sphere.Position().Minus(pos)
		path := s.Forward().Mult(length)
		if rel.Project(path) < s.Radius()+sphere.Radius() {
			distance := rel.Len()
			if distance >= length || distance == 0 { // Too far away, don't care
				continue
			}
			//			fmt.Println("collision potential", sphere.Position())
			if distance < closestDistance {
				//				fmt.Println("closest!", sphere.Position())
				closest = sphere
				closestDistance = distance
			}
		}
	}

	if closest == nil {
		return nil
	}
	//	fmt.Println("fleeing", closest.Position(), "from", s.Position())
	return Flee(s, closest.Position())
}

func Contain(s Steerer, leftTopBack Vector, rightBottomFront Vector, futureUpdates float32, dimensions int) Vector {
	future := s.Position().Plus(s.Forward().Mult(futureUpdates))

	for i := 0; i < dimensions; i++ {
		gradient := Vector{0, 0, 0}
		gradient[i] = 1
		if leftTopBack[i] > future[i] {
			normal := gradient.Normal(future, (future[i]-leftTopBack[i])*futureUpdates)
			fmt.Println("too small", i, gradient, future, normal)

			return Seek(s, normal)
			//			return Seek(s, s.Side().Mult(leftTopBack[i]-future[i]))
		}
		if rightBottomFront[i] < future[i] {

			normal := gradient.Normal(future, (rightBottomFront[i]-future[i])*futureUpdates)
			fmt.Println("too big", i, gradient, future, normal)
			fmt.Println(s.Side().Hadamard(normal))
			return Seek(s, normal)

		}
	}
	return Vector(nil)
}
