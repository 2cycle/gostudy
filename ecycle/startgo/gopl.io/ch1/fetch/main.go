package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)	//httpGet은 http요청을 생성하고 오류가 없으면 응답구조체를 resp로 반환
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)	//resp의 Body에는 응답이 포함돼있다. ,ReadAll같은 경우 전체 응답을 읽는다.
		resp.Body.Close()	//stream resource 관리를 위해 닫아줌
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n",url, err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)

	}
}
