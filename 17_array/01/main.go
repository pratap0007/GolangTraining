package main

import "fmt"

func main() {
	x := [5]int{2, 3, 4, 5, 0}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[4])
	x[4] = 777
	fmt.Println(x[4])
}
