package main

import "fmt"

func main() {

	// Golang doesn't have while loops
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}

	// ARRAY vs SLICES

	prices := [...]int{1, 2, 3, 4, 5} // array
	var prices2 = [3]int{4, 5, 6}
	fmt.Printf("Value = %v, Type = %T\n", prices, prices)                              // Value = [1 2 3 4 5], Type = [5]int
	fmt.Printf("%v Length = %v, Capacity = %v\n", prices2, len(prices2), cap(prices2)) // Length = 3, Capacity = 3

	// 2D array
	var matrix [3][4]int // (row, col)
	matrix[1] = [...]int{40, 41, 42, 43}
	fmt.Println(matrix) // [[0 0 0 0] [40 41 42 43] [0 0 0 0]]

	// arrays are passed by value in golang
	a := [...]int{11, 12, 13, 14, 15}
	b := a
	b[0] = 123
	fmt.Printf("a=%v, b=%v\n", a, b)

	for i, value := range a {
		fmt.Printf("(%v, %v)", i, value)
	}
	fmt.Println()
	for _, value := range b {
		fmt.Printf("%v,",value)
	}

}
