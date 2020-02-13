package main

type Vehicle struct {
	Mass        float32
	Position    Vector
	Velocity    Vector
	MaxForce    float32
	MaxSpeed    float32
	Orientation Vector
}

func (v Vehicle) Update() {

}

func (v Vehicle) Move() {
	steering_direction := Vector{1, 2, 0} // TODO: Set this to something good
	steering_force := steering_direction.Truncate(v.MaxForce)
	acceleration := steering_force.DividedBy(v.Mass)
	v.Velocity = v.Velocity.Plus(acceleration).Truncate(v.MaxSpeed)
	v.Position = v.Position.Plus(v.Velocity)
}
