package main

import "fmt"

var ok = 10

func nested() {

	var ok = 5
	fmt.Println("inside nope:", ok)

}

func main() {
	fmt.Println("inside main:", ok)

	nested()

	fmt.Println("inside main:", ok)
}
