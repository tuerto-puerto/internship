package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello мир, манера крутит мир")
	name := " "
	fmt.Println("скажи мне кто ты")
	fmt.Scanln(&name)
	fmt.Println("hello " + name)
}
