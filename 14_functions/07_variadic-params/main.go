package main

import "fmt"

func main() {
	n := average("a", "b", "c", "d")
	fmt.Println(n)
}

func average(sf ...string) string {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total string
	for _, v := range sf {
		total += v
	}
	return total
}
