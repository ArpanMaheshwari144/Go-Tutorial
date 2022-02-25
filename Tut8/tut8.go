package main

import "fmt"

func main() {

	var name string = "Arpan Maheshwari"
	// fmt.Println(len(name))
	// fmt.Println(name + " is a good boy")

	const pi float64 = 3.1412435
	win := true
	x := 5
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.3f\n", pi)
	fmt.Printf("%T\n", name) // check type
	fmt.Printf("%t\n", win)  // boolean value
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", 25) // binary value
	fmt.Printf("%c\n", 33) // character value
	fmt.Printf("%x\n", 15) // hex code
	fmt.Printf("%e\n", pi) // scientific notations
}
