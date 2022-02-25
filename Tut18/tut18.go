package main

import "fmt"

func main() {

	// firstRun()
	// secondRun()

	defer firstRun()
	secondRun()
}

func firstRun()  { fmt.Println("I Executed First") }
func secondRun() { fmt.Println("I Executed Second") }
