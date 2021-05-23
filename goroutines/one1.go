package main

import (
	"fmt"	
	"net/http"
)


func checkURL(url string, channel chan string){
	resp, err := http.Get(url)

	if err != nil {
		channel <- "website" + url + "is DOWN"
	} else {
		channel <- url + " - " + (*resp).Status
	}
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

	for i := 0; i<len(urls); i++{
		fmt.Println(<-channel)
	}

}