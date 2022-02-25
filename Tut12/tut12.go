package main

import "fmt"

func main() {

	// making of map which contains string(as name) and int(as age)
	studentAge := make(map[string]int)

	studentAge["Arpan"] = 20
	studentAge["Adarsh"] = 20
	studentAge["Verma"] = 20
	studentAge["Patra"] = 20
	studentAge["Akbar"] = 20
	studentAge["Aditya"] = 20

	// fmt.Println(studentAge)
	// fmt.Println(studentAge["Arpan"])
	fmt.Println(len(studentAge))

}
