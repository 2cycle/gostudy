package main

import "fmt"

const boilingF = 212.0 //패키지 선언 수준 main, package 내의 모든 곳에서 사용가능하다.

func main() {
	var f = boilingF //f 변수의 scope는 main 내에서만 한정됨
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)
}