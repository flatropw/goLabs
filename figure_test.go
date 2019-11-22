package main

import (
	"testing"
)

func TestCircle_Area(t *testing.T) {
	given := float64(3)
	var c Figure = Circle{given}
	expected := 28.274333882308138

	t.Run("", func(t *testing.T) {
		got := c.Area()
		if got != expected {
			t.Errorf("got: %v expected: %v given radius: %v", got, expected, given)
		}
	})
}

func TestCircle_Perimeter(t *testing.T) {
	given := float64(3)
	var c Figure = Circle{given}
	expected := 18.84955592153876

	t.Run("", func(t *testing.T) {
		got := c.Perimeter()
		if got != expected {
			t.Errorf("got: %v expected: %v given radius: %v", got, expected, given)
		}
	})
}

func TestSquare_Area(t *testing.T) {
	given := float64(3)
	var s Figure = Square{given}
	expected := float64(9)

	t.Run("", func(t *testing.T) {
		got := s.Area()
		if got != expected {
			t.Errorf("got: %v expected: %v given round: %v", got, expected, given)
		}
	})
}

func TestSquare_Perimeter(t *testing.T) {
	given := float64(3)
	var s Figure = Square{given}
	expected := float64(12)

	t.Run("", func(t *testing.T) {
		got := s.Perimeter()
		if got != expected {
			t.Errorf("got: %v expected: %v given round: %v", got, expected, given)
		}
	})
}
