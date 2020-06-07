package main

import "fmt"

//low level programming
func main() {

	q := 10
	w := q //clone
	fmt.Println(q, w)

	q = 15
	fmt.Println(q, w)

	a := 10
	b := &a // save a address
	fmt.Println(a, b)

	*b = 20
	fmt.Println(a, *b)

	a = 15
	fmt.Println(a, *b)

}
