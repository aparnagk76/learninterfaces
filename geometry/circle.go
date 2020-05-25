package geometry

const pie = 3.14

type Circle struct {
	X, Y   float64
	Radius float64
}

func (s Circle) Area() float64 {
	return pie * s.Radius * s.Radius
}

func (s Circle) Perimeter() float64 {
	return 2 * pie * s.Radius
}
