// Concept of polymorphism in Golang
package main

import (
	"fmt"
)

// Figure ...
type Figure interface {
	Area() float64
}

// Rectangle ...
type Rectangle struct {

	// declaring struct variables
	length float64
	width  float64
}

// Square ...
type Square struct {

	// declaring struct variable
	side float64
}

// Area ...
func (rect Rectangle) Area() float64 {

	// Area of rectangle = l * b
	area := rect.length * rect.width
	return area
}

// Area ...
func (sq Square) Area() float64 {

	// Area of square = a * a
	area := sq.side * sq.side
	return area
}

// main function
func main() {

	// declaring a rectangle instance
	rectangle := Rectangle{

		length: 10.5,
		width:  12.25,
	}

	// declaring a square instance
	square := Square{

		side: 15.0,
	}

	// printing the calculated result
	fmt.Printf("Area of rectangle: %.3f unit sq.\n", rectangle.Area())
	fmt.Printf("Area of square: %.3f unit sq.\n", square.Area())
}
