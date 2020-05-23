package main

import "fmt"

func nope() {
	const ok = true
	var hello = "hello"

	_ = hello
}

func main() {
	fmt.Println(hello, ok) // no inside the scope
}
