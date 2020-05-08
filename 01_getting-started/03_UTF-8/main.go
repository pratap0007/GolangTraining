package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %c \n", i, i, i, i)
		fmt.Printf("%d \t %b \t %x \t %c \n", i, i, i, i)
	}
	// %c is also applicable for  printing charcters

}
