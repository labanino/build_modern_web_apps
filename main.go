package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, someThinElseWasSaid := saySomething()

	fmt.Println("The function returned", whatWasSaid, someThinElseWasSaid)
}

func saySomething() (string, string) {
	return "something", "and goodbye"
}