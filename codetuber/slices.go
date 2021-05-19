package main

import "fmt"

func main() {

	// SLICE have direct reference to underlying array
	// change in 1 slice == change in all other slice of that underlying array

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array, type= [5]int
	b := a[:3]
	fmt.Printf("%v, type=%T cap=%v\n", b, b, cap(b)) // [1 2 3], type=[]int, cap=10

	c := make([]int, 3)
	fmt.Printf("%v, cap=%v\n", c, cap(c))
	c = append(c, 200, 300, 400, 500, 600)
	fmt.Printf("%v, cap=%v\n", c, cap(c)) // type=[]int, cap=8

	d := make([]int, 0, 0)
	d = append(d, 10, 20, 30)
	fmt.Printf("%v, cap=%v\n", d, cap(d)) // [10 20 30], cap=4
	d = append(d, 10, 20, 30)
	fmt.Printf("%v, cap=%v\n", d, cap(d)) // [10 20 30 10 20 30], cap=8

	e := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
	}

	fmt.Println(e)

}
