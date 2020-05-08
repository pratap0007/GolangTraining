package main

import "fmt"

func main() {
	for i := 10; i < 20; i++ {
		fmt.Printf("%d \t %b \t %#X \n", i, i, i)
	}
}
