package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	name string
	age  int
}
type address struct {
	houseno int
	street  string
	state   string
}
type listaddress struct {
	firstaddress address
}

func main() {
	c1 := customer{"Todd", 44}
	p := customer{"shiv", 23}
	add := address{23, "lauwar", "pratapgarh-uttar pradesh"}
	fmt.Println("type", reflect.TypeOf(add))
	display(add)
	fmt.Println(p)
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe(&p)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x8201e4120
}
func display(x address) {
	x.street = "lauwar and tahsil patti"
	fmt.Println("adress", x)
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Todd 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x8201e4120

}
