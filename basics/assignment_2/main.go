package main

import (
	"fmt"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}

func (s square) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}

func printArea(s shape) float64 {
	return s.getArea()
}

func main() {
	t1 := triangle{10, 40}
	fmt.Println("Area of Trigangle is:", printArea(t1))
	s1 := square{10}
	fmt.Println("Area of Square is:", printArea(s1))
}
