package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	height float64
	width  float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func main() {
	c1 := circle{3.14}
	r1 := rect{2.5, 3.5}

	s1 := []shape{c1, r1}

	fmt.Printf("%+v type=%T\n", s1, s1) // [{radius:3.14} {height:2.5 width:3.5}] type=[]main.shape

	c2 := s1[0].radius  // ERROR : s1[0].radius undefined (type shape has no field or method radius)
	fmt.Println(c2)

	c3 := s1[0].(circle)
	radius := c3.radius
	fmt.Printf("%+v type=%T %v", c3, c3, radius)

}
