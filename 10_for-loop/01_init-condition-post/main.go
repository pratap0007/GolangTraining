package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}
	for j := 1; j <= 5; j = j + 2 {
		fmt.Println(j)
	}
}
