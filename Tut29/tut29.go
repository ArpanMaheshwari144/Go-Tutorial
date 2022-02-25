// Concept of multiple inheritances in Golang
package main

import (
	"fmt"
)

// First ...
type First struct {

	// declaring struct variable
	baseOne string
}

// Second ...
type Second struct {

	// declaring struct variable
	baseTwo string
}

// function to return
// first struct variable
func (f First) printBase1() string {

	// returns a string
	// of first struct
	return f.baseOne
}

// function to return
// second struct variable
func (s Second) printBase2() string {

	// returns a string
	// of first struct
	return s.baseTwo
}

// Child ...
type Child struct {

	// anonymous fields,
	// struct embedding
	// of multiple structs
	First
	Second
}

// main function
func main() {

	// declaring an instance
	// of child struct
	c1 := Child{

		// child struct can directly
		// access base struct variables
		First{
			baseOne: "In base struct 1.",
		},
		Second{
			baseTwo: "\nIn base struct 2.\n",
		},
	}

	// child struct can directly
	// access base struct methods

	// printing base method
	// using instance of child struct
	fmt.Println(c1.printBase1())
	fmt.Println(c1.printBase2())
}
