// Concept of encapsulation using exported function
package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Creating a slice of strings
	slc := []string{"Arpan", "Adarsh", "Verma"}

	// Convert the case of the elements of the given slice using ToUpper() function
	for x := 0; x < len(slc); x++ {

		// Exported Method
		res := strings.ToUpper(slc[x])

		// Exported Method
		fmt.Print(res + " ")
	}
}
