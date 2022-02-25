package main

import (
	"fmt"
)

// This function handles the panic occur in entry function with the help of the recover function
func handlepanic() {

	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

// Function
func entry(lang *string, aname *string) {

	// Deferred function
	defer handlepanic()

	// When the value of lang is nil it will panic
	if lang == nil {
		panic("Error: Language cannot be nil")
	}

	// When the value of aname is nil it will panic
	if aname == nil {
		panic("Error: Author name cannot be nil")
	}
	fmt.Printf("Author Language: %s\nAuthor Name: %s\n", *lang, *aname)
	fmt.Printf("Return successfully from the entry function")
}

// Main function
func main() {

	lang := "GO Language"
	// author := "Arpan"
	entry(&lang, nil)
	fmt.Printf("Return successfully from the main function")
}
