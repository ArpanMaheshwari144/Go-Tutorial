package main

import "fmt"

func main() {

	fmt.Print("Enter First Number: ")
	var first int
	fmt.Scanln(&first)

	fmt.Print("Enter Second Number: ")
	var second int
	fmt.Scanln(&second)

	fmt.Printf("The addition of %d and %d is ", first, second)
	fmt.Print(add(first, second))
}

func add(first, second int) int {
	return first + second
}
