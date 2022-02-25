package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Creating a file
	file, err := os.Create("Tut24/sample.txt")

	if err != nil {

		log.Fatal(err)

	}
	file.WriteString("Arpan Maheshwari")
	file.Close()

	// Reading a file
	stream, err := ioutil.ReadFile("Tut24/sample.txt")

	if err != nil {

		log.Fatal(err)

	}
	str := string(stream)
	fmt.Println(str)
}
