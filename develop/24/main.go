package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	d := p1.Distance(p2)
	fmt.Printf("distance beetween points p1 and p2 is %f \n", d)

	p3 := NewPoint(1, 2)
	p4 := NewPoint(3, 3)
	d2 := p3.Distance(p4)
	fmt.Printf("distance beetween points p3 and p4 is %f \n", d2)
}

type Point struct {
	x,
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
