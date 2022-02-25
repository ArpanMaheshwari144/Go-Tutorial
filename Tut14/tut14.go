package main

import "fmt"

func main() {

	x, y := 5, 6

	fmt.Println(add(x, y))
}

// func functionname(argument-1, argument-2,........,argument-n type-of-argument) return-type {
// 	function-body
// }

func add(x, y int) int {
	return x + y
}
