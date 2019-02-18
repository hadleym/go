package main

import "fmt"
type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

func main() {

	sq := square {
		sideLength: 1,
	}

	tr := triangle {
		base: 1,
		height: 4,
	}

	printArea(sq)
	printArea(tr
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area of shape: ", s.getArea())
}