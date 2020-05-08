package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	meet()
}
func meet() {
	fmt.Println("meet")
}
