package main

import (
	"log"
)

func main() {
	// var mySlice []int

	// mySlice = append(mySlice, "Paola")
	// mySlice = append(mySlice, "Ivan")
	// mySlice = append(mySlice, "Lucas")

	// mySlice = append(mySlice, 3)
	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 1)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:6])
}

// https://golangdocs.com/slices-in-golang