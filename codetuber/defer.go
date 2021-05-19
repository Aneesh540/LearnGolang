package main

import (
	"fmt"
)

func main() {
	// defer keyword is executed at the end;
	// when function is about to return
	// pushed to defer stack with current value
	// defer IS FUNCTIONAL SCOPE NOT BLOCK SCOPE

	var a int = 12
	defer fmt.Println(a) // print 12 at the end
	a = 344
	fmt.Println(a)

	f := createFile("./defer.txt")
	defer closeFile(f) // so we didn't forgot to close file
	writeFile(f)
}
