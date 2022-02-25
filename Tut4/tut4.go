package main

import "fmt"

func main() {

	x, y := 6, 6

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x mod y = ", x%y)

	isbool1 := true
	isbool2 := false

	fmt.Println(isbool1 && isbool2)
	fmt.Println(isbool1 || isbool2)
	fmt.Println(!isbool1)
	fmt.Println(!isbool2)

}
