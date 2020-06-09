package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 각 요청은 핸들러를 호출한다
	log.Fatal(http.ListenAndServe(":8000", nil)) //8000번 포트로 들어오는 요청을 처리하는 서버를 시작
}

// handler는 요청된 URL request의 Path 구성 요소를 반환한다.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

/*
1. 요청은 http.Request구조체로 표현되며 그 안에는 들어온 요청의 URL을 비롯한 다수의 관련 필드가 있다.
	r.Body, r.Header, r.Host, r.Form
2. 요청이 도착 > 핸들러함수로 넘김 > 응답돌려줌 (fmt.Fprintf)


 */

