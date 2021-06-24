package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFail = errors.New("Request fail")

type result struct {
	url    string
	status string
}

func main() {
	// results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i, _ := range urls {
		fmt.Println("Received number : ", i, " result: ", <-c)
	}
}

// if the parameter is looks like chan<-, it can only send
// and the parameter looks like chan->, it can only receive
func hitURL(url string, c chan<- result) {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, status: "FAILED"}
	} else {
		c <- result{url: url, status: "SUCCESS"}
	}
}
