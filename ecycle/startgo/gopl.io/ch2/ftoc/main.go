// Ftoc는 화씨-섭씨 변환을 두 번 출력한다.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {	//fToC함수는 온도변환 로직을 캡슐화(encapsulation)하고 있어 한 번 정의한 후 여러 곳에서 사용
	return (f - 32) * 5 / 9
}