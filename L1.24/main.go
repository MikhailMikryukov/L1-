package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := other.GetX() - p.GetX()
	dy := other.GetY() - p.GetY()
	return math.Sqrt(dx*dx + dy*dy)
}

// L1.24 Расстояние между точками
func main() {
	point1 := NewPoint(2.0, 2.0)
	point2 := NewPoint(4.0, 4.0)

	fmt.Println(point1.Distance(point2))

}
