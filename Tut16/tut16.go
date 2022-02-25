package main

import "fmt"

func main() {

	fmt.Print("Enter Your First Name: ")
	var first string
	fmt.Scanln(&first)

	fmt.Print("Enter Your Last Name: ")
	var second string
	fmt.Scanln(&second)

	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second)
}
