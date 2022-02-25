// Concept of inheritance in Golang
package main

import (
	"fmt"
)

// Comic ...
type Comic struct {

	// declaring struct variable
	Universe string
}

// ComicUniverse ...
func (comic Comic) ComicUniverse() string {

	// returns comic universe
	return comic.Universe
}

// Marvel ...
type Marvel struct {

	// anonymous field,
	// this is composition where
	// the struct is embedded
	Comic
}

// DC ...
type DC struct {

	// anonymous field
	Comic
}

// main function
func main() {

	// creating an instance
	c1 := Marvel{

		// child struct can directly
		// access base struct variables
		Comic{
			Universe: "MCU",
		},
	}

	// child struct can directly
	// access base struct methods

	// printing base method using child
	fmt.Println("Universe is:", c1.ComicUniverse())

	c2 := DC{
		Comic{
			Universe: "DC",
		},
	}

	// printing base method using child
	fmt.Println("Universe is:", c2.ComicUniverse())
}
