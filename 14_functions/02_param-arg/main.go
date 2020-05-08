package main

import "fmt"

func main() {
	greet("Jane", 48)
	greet("John", 45)
}

func greet(name string, x int) {
	fmt.Println(name, x)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
