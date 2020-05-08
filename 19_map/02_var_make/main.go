package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	var x = make(map[string]int)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	x["shu"] = 32632
	fmt.Print(x)

	fmt.Println(myGreeting)
}
