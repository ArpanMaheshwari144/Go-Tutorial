package main

import "fmt"

func main() {

	fmt.Println(div(3, 0))
	myPanic()
}

func div(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2

	return solution
}

func myPanic() {

	defer func() {
		fmt.Println(recover())
	}()
	panic("Panic")
}
