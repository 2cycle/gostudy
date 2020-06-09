package main

import "fmt"

func main() {
	sliceA := []int{0, 1, 2}
	sliceB := make([]int, len(sliceA), cap(sliceA)*2) //sliceA에 2배 용량인 슬라이스 선언

	copy(sliceB, sliceA) //A를 B에 붙여넣는다

	fmt.Println(sliceB)                   // [0 1 2 ] 출력
	fmt.Println(len(sliceB), cap(sliceB)) // 3, 6 출력

	sliceA[0] = 10
	fmt.Println(sliceA, sliceB)

	println(&sliceA, &sliceB)

	/*
		-------- 실행 결과
		[0 1 2]
		3 6
		0xc000104f60 0xc000104f48	// 새로운 메모리가 할당됨. copy()
	*/

	fmt.Println()
	sliceC := make([]int, 0, 3) //용량이 3이고 길이가0인 정수형 슬라이스 선언
	sliceC = append(sliceC, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sliceC), cap(sliceC))

	sliceD := sliceC[1:3] //인덱스 1요소부터 2요소까지 복사
	fmt.Println(sliceD)

	sliceD = sliceC[2:] //인덱스 2요소부터 끝까지 복사
	fmt.Println(sliceD)

	sliceD[0] = 6

	fmt.Println(sliceC) //슬라이스 l의 값을 바꿨는데 c의 값도 바뀜
	//값을 복사해온 것이 아니라 기존 슬라이스 주솟값을 참조

	a := []int{1, 200, 3000, 40000}
	b := a[0:3] // [1 200 3000] : 0번 인덱스부터 3번 인덱스 전까지 Slicing 합니다.
	c := b[1:3] // [200 3000]

	b[1] = 0 //
	fmt.Println(a, b, c)
}
