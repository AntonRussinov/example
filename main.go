package main

import (
	"fmt"
)

func main() {

	go fmt.Println("1!")
	go fmt.Println("2!")
	go fmt.Println("3!")
	go fmt.Println("42112!")

	fmt.Scanln()
}
