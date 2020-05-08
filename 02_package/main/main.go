package main

import (
	"fmt"
	someAlias "github.com/pratap0007/Golang-Basic/GolangTraining/02_package/icomefromalaska"
	// "github.com/pratap0007/Golang-Basic/GolangTraining/02_package/icomefromalaska"
	"github.com/pratap0007/Golang-Basic/GolangTraining/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(someAlias.Fullname)
}
