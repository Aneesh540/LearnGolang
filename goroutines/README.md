# SUMMARY

## 1. Function to check site is down or not
> **FILENAME - one1.go**
```golang
func checkURL(url string, channel chan string){
	resp, err := http.Get(url)
	if err != nil {
  //
		channel <- "website" + url + "is DOWN"
	} else {
		channel <- url + " - " + (*resp).Status
	}
}
```
#### Basic go-routine code (CORRECT)
```golang
func main() {
	urls := []string{"http://facebook.com","http://golang.org/werer"}

	channel := make(chan string)

	for _, url := range urls {
		go checkURL(url, channel)
	}
	for i := 0; i<len(urls); i++{
		fmt.Println(<-channel)
	}

}
```

## 2. Using Enhanced for loop for channels(WRONG)
> **FILENAME - two2.go**
> 
> "link" variable is shared in main & child goroutine, which is wrong (RACE CONDITION)
> 
> check status in every 5 seconds
```golang
	channel := make(chan string)

	for _, url := range urls {
		go checkURL(url, channel)
	}
	for link := range channel{
		go func(){
			time.Sleep(5* time.Second)
			checkURL(link, channel)
		}()
	}
```
![GITHUB_GOLANG](https://user-images.githubusercontent.com/25201571/119268624-7185eb80-bc11-11eb-9794-fa76554f307e.PNG)


## 3. Using enhanced for loop (CORRECT)
> **FILENAME - three.go**
>
> As its a function, so all argument will be passed by value. Hence "link" variable is created for that function
```golang
	for link := range channel{
		go func(l string){
			time.Sleep(5* time.Second)
			checkURL(l, channel)
		}(link)
	}
 ```

## MUST CHECK LINKS
> [https://www.youtube.com/watch?v=LvgVSSpwND8](https://www.youtube.com/watch?v=LvgVSSpwND8) Advance topic on goroutine
