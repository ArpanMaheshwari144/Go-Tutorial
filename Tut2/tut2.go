package main

import "fmt"

func main() {

	// a constant is a variable with a value that cannot be changed
	const pi float64 = 3.141592265359
	fmt.Println(pi)

	// we can declare multiple varaibles like this
	var (
		varA = 2
		varB = 3
	)
	fmt.Println(varA, varB)

	// strings are a series of characters between "" or ''
	var Name string = "Arpan Maheshwari"

	fmt.Println(len(Name))
}
