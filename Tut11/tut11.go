package main

import "fmt"

func main() {

	// var studentCount [10]int  // Array of 10 size
	// for i := 0; i < 10; i++ {
	// 	studentCount[i] = i + 1
	// 	fmt.Println(studentCount[i])
	// }

	// var evenNum [5]int // Array of 5 size
	// evenNum[0] = 0
	// evenNum[1] = 2
	// evenNum[2] = 4
	// evenNum[3] = 6
	// evenNum[4] = 8
	// fmt.Println(evenNum[0])
	// fmt.Println(evenNum[1])
	// fmt.Println(evenNum[2])
	// fmt.Println(evenNum[3])
	// fmt.Println(evenNum[4])

	// evenNum := [5]int{0, 2, 4, 6, 8} // Shorthand making of an Array of 5 size
	// fmt.Println(evenNum[0])
	// fmt.Println(evenNum[1])
	// fmt.Println(evenNum[2])
	// fmt.Println(evenNum[3])
	// fmt.Println(evenNum[4])

	// evenNum := [5]int{0, 2, 4, 6, 8}

	// for-each loop
	// for _, value := range evenNum {
	// 	fmt.Println(value)
	// }

	// for index, value := range evenNum {
	// 	fmt.Println(index, "->", value)
	// }

	// numSlice := []int{5, 4, 3, 2, 1}
	// sliced := numSlice[3:5]
	// fmt.Println(sliced)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5, 10)
	copy(slice2, slice1)
	fmt.Println(slice2)
	fmt.Println(slice1)

	slice3 := append(slice1, 6, 7)
	fmt.Println(slice3)
}
