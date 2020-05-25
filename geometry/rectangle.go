package geometry

type Rectangle struct {
	X float64
	Y float64
	L float64
	B float64
}

func (s Rectangle) Area() float64 {
	return s.L * s.B
}
func (s Rectangle) Perimeter() float64 {
	return 2 * (s.L + s.B)
}
