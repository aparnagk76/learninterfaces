package main

import (
	"fmt"

	drw "github.com/aparnagk76/learninterfaces/drawing"
	geo "github.com/aparnagk76/learninterfaces/geometry"
)

func main() {
	var arrayOfShapes = []geo.Shape{
		geo.Traingle{X1: 12.0, X2: 8.0, X3: 6.0, Y1: 12.0, Y2: 8.0, Y3: 10.0},
		geo.Square{X: 10, Y: 20, Side: 5},
		geo.Circle{Radius: 5.0, X: 12.0, Y: 8.0},
	}

	drawing := drw.Drawing{Shapes: arrayOfShapes}
	fmt.Println(drawing.NetArea())

	//fmt.Println(computeArea(geo.Traingle{X1: 12.0, X2: 8.0, X3: 6.0, Y1: 12.0, Y2: 8.0, Y3: 10.0}))
	//fmt.Println(computeArea(geo.Traingle{X1: 12.0, X2: 8.0, X3: 6.0, Y1: 12.0, Y2: 8.0, Y3: 10.0}, geo.Square{X: 10, Y: 20, Side: 5}))
	fmt.Println(computeArea(geo.Traingle{X1: 12.0, X2: 8.0, X3: 6.0, Y1: 12.0, Y2: 8.0, Y3: 10.0}, geo.Square{X: 10, Y: 20, Side: 5}, geo.Circle{Radius: 5.0, X: 12.0, Y: 8.0}))
}

func computeArea(shapes ...geo.Shape) float64 {
	totalArea := 0.0
	for _, value := range shapes {
		totalArea = totalArea + value.Area()
	}
	return totalArea
}
