package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) // To get reference to a variable you use &nameOfTheVariable
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) { // To reference a pointer you use *name
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

// https://golangdocs.com/pointers-in-golang