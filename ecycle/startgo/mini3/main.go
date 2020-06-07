package main

import (
	"errors"
	"fmt"
	"net/http"
)

type resultReqsuet struct {
	url    string
	status string
}

var errRequestFailed = errors.New("requset failed")

func main() {
	//var results = map[string]string {}
	var results = make(map[string]string) //initialize.
	c := make(chan resultReqsuet)

	urls := []string{
		"https://www.daum.net/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.naver.com/",
		"https://www.clien.net/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://melon.com/",
	}

	for _, url := range urls {
		go urlcheck(url, c)

	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}

func urlcheck(url string, c chan<- resultReqsuet) { //cahnnel send only
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- resultReqsuet{url: url, status: status}
}
