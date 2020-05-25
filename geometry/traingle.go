package geometry

import "math"

type Traingle struct {
	X1, Y1 float64
	X2, Y2 float64
	X3, Y3 float64
}

func (s Traingle) Area() float64 {
	a, b, c := s.GetAllSides()
	semiPerimeter := s.Perimeter() / 2
	return math.Sqrt(semiPerimeter * (semiPerimeter - a) * (semiPerimeter - b) * (semiPerimeter - c))

}
func (s Traingle) Perimeter() float64 {
	a, b, c := s.GetAllSides()
	return a + b + c
}

func (s Traingle) GetAllSides() (float64, float64, float64) {
	a := math.Sqrt(math.Pow(s.X2-s.X1, 2) + math.Pow(s.Y2-s.Y1, 2))
	b := math.Sqrt(math.Pow(s.X3-s.X2, 2) + math.Pow(s.Y3-s.Y2, 2))
	c := math.Sqrt(math.Pow(s.X3-s.X1, 2) + math.Pow(s.Y3-s.Y1, 2))
	return a, b, c
}
