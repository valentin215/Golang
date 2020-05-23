package main

import "fmt"

func main() {
	fmt.Printf("Hello, World\n")
	fmt.Println("Mehdi")

	bye()
	hey()
}

func hey() {
	fmt.Println("Hey!")
}

func bye() {
	fmt.Println("Bye")
}
