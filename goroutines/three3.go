package main

import (
	"fmt"
	"net/http"
	"time"
)


func checkURL(url string, channel chan string){
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("website" + url + "is DOWN")
	} else {
		fmt.Println(url + " - " + (*resp).Status)
	}

	channel <- url
}

func main() {
	urls := []string{
		"http://facebook.com",
		"http://golang.org/werer",
		"http://google.com",
		"http://linkedin.com",
	}

	channel := make(chan string)

	for _, url := range urls {
		go checkURL(url, channel)
	}

	for link := range channel{
		go func(l string){
			time.Sleep(5* time.Second)
			checkURL(l, channel)
		}(link)
	}


}