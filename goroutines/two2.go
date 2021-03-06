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
		go func(){
			time.Sleep(5* time.Second)
			checkURL(link, channel) 
			// since link is not a function variable of this new go routine SO
			// its value will be taken(referenced) from parent's go routine function
			// (referenced) = RACE CONDITION
		}()
	}


}