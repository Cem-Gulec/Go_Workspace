package main

import (
	"fmt"
	"math"
)

// Defining interface
// interfaces are the method signatures for struct
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// area method for circle
// (value receiver func)
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// area method for circle
// (value receiver func)
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 20, height: 10}

	fmt.Printf("Circle area: %f\nRectangle area: %f\n", circle.area(), rectangle.area())

	// Alternatively
	fmt.Printf("Circle area: %f\nRectangle area: %f", getArea(circle), getArea(rectangle))
}
