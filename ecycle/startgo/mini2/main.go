package main

import (
	"fmt"

	"github.com/ecycle/startgo/mini2/mydict"
)

func main() {

	dictionary := mydict.Dictionary{"first": "first word"}

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	definition1, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition1)
	}
	dictionary.Update("first", "changed first value")
	value, _ := dictionary.Search("first")
	fmt.Println(value)

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}
