package main

import (
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	roundLength float64
}

func (s Square) Perimeter() float64 {
	return s.roundLength * 4
}

func (s Square) Area() float64 {
	return s.roundLength * s.roundLength
}

type Circle struct {
	radiusLength float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radiusLength
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radiusLength, 2)
}
