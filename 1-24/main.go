package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow((p.X-other.X), 2) + math.Pow((p.Y-other.Y), 2))
}

func main() {
	p1 := NewPoint(4, 5)
	p2 := NewPoint(5, 6)
	fmt.Println(p2.Distance(p1))
}
