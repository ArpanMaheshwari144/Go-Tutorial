package main

import "fmt"

func main() {

	// var age int = 18
	age := 18

	// if age > 18 {
	// 	fmt.Println("You can vote!!")
	// } else {
	// 	fmt.Println("You cannot vote!!")
	// }

	// if age > 18 {
	// 	fmt.Println("You can vote")
	// } else if age == 18 {
	// 	fmt.Println("wait for 1 year")
	// } else {
	// 	fmt.Println("You cannot vote")
	// }

	switch age {
	case 16:
		fmt.Println("Prepare for college")
	case 18:
		fmt.Println("Going to college")
	case 20:
		fmt.Println("Prepare for a job")
	default:
		fmt.Println("Enjoy your life")

	}
}
