package geometry

type Square struct {
	x    float64
	y    float64
	side float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (s *Square) perimter() float64 {
	return 4 * s.side
}
