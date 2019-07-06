package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sidelength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	t1 := triangle{5.0, 5.0}
	s1 := square{5.0}
	printArea(t1)
	printArea(s1)
}
func printArea(s shape) {
	fmt.Println("The area is ", s.getArea())

}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}
