package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//only primitive type - use struct
	temp := map[string]string{"fname": "wally", "lname": "where"}
	fmt.Println(temp)

	favFood := []string{"sushi", "pizza"}
	wally := person{"wally", 32, favFood}
	wally2 := person{name: "wally", age: 32, favFood: favFood}

	fmt.Println(wally.age, wally.name, wally.favFood)
	fmt.Println(wally2.age, wally2.name, wally2.favFood)
}
