package main

import "fmt"

const GoCon1 int = 12     // starting with capital, it will be exported entity
const goCon2 float64 = 15 // local to this file

const goCon3 = 3 // untyped constant, takes the type as needed by usage

func main() {
	fmt.Printf("value=%v\n", goCon2)
	const goCon2 = "aneesh"
	fmt.Printf("value=%v\n", goCon2) // shadowing of constants

	// additions of constants
	const a = 3.14
	fmt.Println(GoCon1 + a) // error : constant 3.14 truncated to integer
	fmt.Println(goCon3 + a) // this will work
}
