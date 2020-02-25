package steering

type Steerer interface {
	Position() Vector
	Velocity() Vector
	MaxSpeed() float32
	Forward() Vector
	Side() Vector
	Up() Vector
	Radius() float32
}
