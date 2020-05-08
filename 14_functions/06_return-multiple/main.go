package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

func greet(fname, lname string) (p string, q string) {
	p = fmt.Sprint(fname, lname)
	q = fmt.Sprint(lname, fname)
	return
}
