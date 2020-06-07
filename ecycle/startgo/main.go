package main

//go has so many lib for use
import (
	"fmt"
	"strings"
)

/*

main package used only compiling - (not lib, opensource)
compiler find main package at first and compile
*/

func multiply(a int, b int) int {
	return a * b
}

func multiply2(a, b int) int {
	return a * b
}

//multireturn
func lenghtAndUpper(chars string) (int, string) {
	return len(chars), strings.ToUpper(chars)
}

//return, defer - when the function finished this started
func lenghtAndUpper2(chars string) (lengths int, uppers string) {
	//this will used for api call, stream close, after function
	defer fmt.Println("lengthAndUpeer2 is finished")
	lengths = len(chars)
	uppers = strings.ToUpper(chars)
	return
}

//repeat
func repeat(words ...string) {
	fmt.Println(words)
}

//loop
func superAdd(nums ...int) int {
	//range is loop for array
	fmt.Println(nums)

	/*
		for i:=0; i<len(nums); i++ {
			fmt.Println(nums[i])
		}
	*/

	total := 0
	// range return index and number
	for index, num := range nums {
		fmt.Println(index, num)
		total += num
	}

	return total
}

//switch and variable expression
func isKoreaAdultAge(age int) bool {
	//if koreanAge := age + 2; koreanAge < 18 {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

func main() {
	fmt.Println("Hello world")

	var name1 string = "wally"
	name2 := "wally"
	//this can only inside of the function
	fmt.Println(name1 + name2)

	//function test
	num1 := 5
	num2 := 10

	fmt.Println(multiply(num1, num2))
	fmt.Println(multiply2(num1, num2))

	//multyreturn test
	chars := "wally.where"
	lenChars, upperChars := lenghtAndUpper(chars)
	lenChars1, _ := lenghtAndUpper(chars)
	lenChars2, upperChars2 := lenghtAndUpper2(chars)
	// _ is pass(ignored)
	fmt.Println(lenChars, upperChars)
	fmt.Println(lenChars1)
	fmt.Println(lenChars2, upperChars2)

	//repeat test
	repeat("wally1", "wally2", "wally3", "wally4", "wally5")

	//superAdd
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("total is %d\n", total)

	//switch and variable expression
	fmt.Println(isKoreaAdultAge(18))
	fmt.Println(isKoreaAdultAge(17))

}
