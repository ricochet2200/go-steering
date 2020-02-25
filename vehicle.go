package main

type Vehicle struct {
	Mass        float32
	position    Vector
	velocity    Vector
	MaxForce    float32
	maxSpeed    float32
	Orientation Matrix
	radius      float32
}

func MakeVehicle(pos Vector, mass, maxForce, maxSpeed, radius float32) *Vehicle {
	v := Vector{1, 1, 1}
	data := []float32{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	o := MakeMatrix(3, 3, data)
	return &Vehicle{mass, pos, v, maxForce, maxSpeed, o, radius}
}

func (v *Vehicle) X() float32 {
	return v.position[0] + v.Radius()/2
}

func (v *Vehicle) Y() float32 {
	return v.position[1] + v.Radius()/2
}
func (v *Vehicle) Position() Vector {
	return v.position.Clone()
}
func (v *Vehicle) Velocity() Vector {
	return v.velocity.Clone()
}
func (v *Vehicle) MaxSpeed() float32 {
	return v.maxSpeed
}

func (v *Vehicle) Radius() float32 {
	return v.radius
}

func (v *Vehicle) Forward() Vector {
	return v.Orientation[0]
}

func (v *Vehicle) Side() Vector {
	return v.Orientation[1]
}

func (v *Vehicle) Up() Vector {
	return v.Orientation[2]
}

func (v *Vehicle) Move(steering_direction Vector) {
	//	fmt.Println("steering direction", steering_direction)
	steering_direction[2] = 0

	steering_force := steering_direction.Trunc(v.MaxForce)
	//	fmt.Println("force", steering_force)
	acceleration := steering_force.DividedBy(v.Mass)
	//	fmt.Println("acceleration", acceleration)
	v.velocity = v.velocity.Plus(acceleration).Trunc(v.maxSpeed)
	//	fmt.Println("velocity", v.velocity)
	v.position = v.position.Plus(v.velocity)
	//	fmt.Println("position", v.position)

	v.position[2] = 0 // We are only in 2D, so keep this from wandering

	new_forward := v.velocity.Normalize()
	approx_up := v.Orientation[2]
	new_side := new_forward.Cross(approx_up)
	new_up := new_forward.Cross(new_side)

	v.Orientation[0] = new_forward
	v.Orientation[1] = new_side
	v.Orientation[2] = new_up
}
