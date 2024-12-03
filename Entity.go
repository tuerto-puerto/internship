package main

import "fmt"

type User struct {
	id        int
	firstName string
	lastName  string
	age       string
	email     string
	password  string
}

func main() {
	user := User{
		id:        1,
		firstName: "пальто",
		lastName:  "пацан",
		age:       "18 лет",
		email:     "пальто.пацан@gmail.com",
		password:  "123	",
	}

	fmt.Println("User", user)
}
