package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	/* some hardcoded stuff
	i used the minus sign in calculation of y axis because in the attached screenshot start.y lower then end.y (???)
	*/
	return Point{s.start.x + int(s.a), s.start.y - int(s.a)}
}

func (s Square) Perimeter() uint {
	const squareEdgesCount = 4
	return s.a * squareEdgesCount
}

func (s Square) Area() uint {
	return uint(math.Pow(float64(s.a), float64(2)))
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
