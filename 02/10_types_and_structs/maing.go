package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Alex",
		LastName:  "Labanino",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

// https://golangdocs.com/structs-in-golang 