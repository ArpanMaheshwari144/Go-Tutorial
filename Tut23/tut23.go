package main

import (
	"fmt"
	"math"
)

func main() {

	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of rectangle is", getArea(rect))
	fmt.Println("Area of circle is", getArea(circ))
}

// Shape is ...
type Shape interface {
	area() float64
}

// Rectangle is ...
type Rectangle struct {
	height float64
	width  float64
}

// Circle is ...
type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {

	return r1.height * r1.width

}

func (c1 Circle) area() float64 {

	// return math.Pi * c1.radius * c1.radius
	return math.Pi * math.Pow(c1.radius, 2)

}

func getArea(shape Shape) float64 {

	return shape.area()

}
