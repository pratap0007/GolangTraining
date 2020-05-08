package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe", 34.89))
}

func greet(fname string, lname string, x float64) (s string) {
	s = fmt.Sprint(fname, lname, x)
	return
}

/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
