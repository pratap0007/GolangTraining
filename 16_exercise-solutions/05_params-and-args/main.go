package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	x := []int{3, 4, 5, 66, 7}
	foo(x...)
	foo()
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
