package drawing

import (
	geo "github.com/aparnagk76/learninterfaces/geometry"
)

type Drawing struct {
	Shapes []geo.Shape
}

func (d Drawing) NetArea() float64 {
	totalValue := 0.0
	for _, value := range d.Shapes {
		totalValue = totalValue + value.Area()
	}
	return totalValue
}
