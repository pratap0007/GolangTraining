package main

import "fmt"

func main() {
	p := map[int]int{}
	myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	p[23] = 2653

	fmt.Println(myGreeting)
	fmt.Println(p)
}
