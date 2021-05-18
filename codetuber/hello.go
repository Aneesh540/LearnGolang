package main

import "fmt"

var global string = "GLOBAL STRING"

func main() {

	// VARIABLES DECLERATIONS
	var str string = "aneesh jain"
	var integer int = 123
	var float float64 = 3.14
	boolean := true
	naman := 11.12
	fmt.Println(integer, str, float, boolean, naman)

	if true {
		global := 123
		fmt.Println(global) // 123
	}
	fmt.Println(global) // GLOBAL STRING

	// NEW WITH STRINGS
	// STRINGS ARE IMMUTABLE
	str2 := "aneesh jain"
	fmt.Println(str2[0], string(str2[0])) // 97 a
	fmt.Printf("Value = %v, Type = %T\n", str2, str2)

	// string to byte slice
	bs := []byte(str2)
	fmt.Printf("Value = %v, Type = %T\n", bs, bs)

}
