package main

import "fmt"

func main() {

	fmt.Println(add(10, 20, 30, 40, 50))
}

func add(args ...int) int {
	sum := 0

	for _, value := range args {
		sum += value
	}
	return sum
}
