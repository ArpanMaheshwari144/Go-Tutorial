// Concept of polymorphism in Golang
package main

import (
	"fmt"
)

// Reading ...
type Reading interface {

	// declaring interface method
	readingTime() int
}

// Book ...
type Book struct {

	// declaring struct variables
	name      string
	pageCount int
}

// Newspaper ...
type Newspaper struct {

	// declaring struct variables
	name      string
	pageCount int
}

// Magazine ...
type Magazine struct {

	// declaring struct variables
	name      string
	pageCount int
}

// function to calculate reading
// time for book
func (book Book) readingTime() int {

	// taking average speed
	// of 10 mins per page
	readTime := 10 * book.pageCount
	return readTime
}

// function to calculate reading
// time for newspaper
func (newspaper Newspaper) readingTime() int {

	// taking average speed
	// of 30 mins per page
	readTime := 30 * newspaper.pageCount
	return readTime
}

// function to calculate reading
// time for magazine
func (magazine Magazine) readingTime() int {

	// taking average speed
	// of 5 mins per page
	readTime := 5 * magazine.pageCount
	return readTime
}

// function to calculate reading time
func calcReadingTime(ReadingTime []Reading) int {

	totalTime := 0

	// looping through elements
	// of the Reading array
	for _, t := range ReadingTime {

		// run time polymophism, call to
		// method depends on object being
		// referred at run time
		totalTime += t.readingTime()
	}

	return totalTime
}

// main function
func main() {

	// declaring a book instance
	book1 := Book{
		name:      "Goosebumps",
		pageCount: 150,
	}

	// declaring a newspaper instance
	newspaper1 := Newspaper{
		name:      "TOI",
		pageCount: 12,
	}

	// declaring a magazine instance
	magazine1 := Magazine{
		name:      "Forbes",
		pageCount: 40,
	}

	// array of type Reading interface
	ReadingTime := []Reading{book1, newspaper1, magazine1}

	// total reading time calculated
	totalTime := calcReadingTime(ReadingTime)

	// Printing total time for reading
	fmt.Printf("Total Time is %d minutes.\n", totalTime)
}
