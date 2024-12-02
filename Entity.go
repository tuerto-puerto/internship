package main

import "fmt"

type User struct {
	id       int
	name     string
	email    string
	password string
}

func main() {
	user := User{
		id:       1,
		name:     "пальто",
		email:    "пальто.пацан@gmail.com",
		password: "123",
	}

	fmt.Println("User", user)
}
