package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSlices = [][]int{
	{1, 3, 5, 8},
	{2, 4, 6, 8},
	{33, 16, 18, 25},
	{100, 100, 100, 100},
	{},
}

func TestCalculateSumOfSlice(t *testing.T) {
	var expectedResults = []int{17, 20, 92, 400, 0}
	for k, v := range testSlices {
		t.Run(fmt.Sprintf("%d:%v", k, v), func(t *testing.T) {
			got := CalculateSumOfSlice(v)
			if got != expectedResults[k] {
				t.Errorf("got: %d expected: %d given: %v", got, expectedResults[k], v)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	var expectedResults = []float64{4.25, 5, 23, 100, 0}
	for k, v := range testSlices {
		t.Run(fmt.Sprintf("%d:%v", k, v), func(t *testing.T) {
			got := Average(v)
			if got != expectedResults[k] {
				t.Errorf("got: %e expected: %e given: %v", got, expectedResults[k], v)
			}
		})
	}
}

var testLenSlices = [][]string{
	{"one", "two", "three", "four"},
	{"bed", "table", "armchair"},
	{"zebra", "elephant", "cat", "rhino"},
	{"mouse", "keyboard", "trackpad", "touchbar"},
	{},
}

func TestMax(t *testing.T) {
	var expectedResults = []string{"three", "armchair", "elephant", "keyboard", ""}
	for k, v := range testLenSlices {
		t.Run(fmt.Sprintf("%d:%v", k, v), func(t *testing.T) {
			got := Max(testLenSlices[k])
			if got != expectedResults[k] {
				t.Errorf("got: %s expected: %s given: %v", got, expectedResults[k], v)
			}
		})
	}
}

var testReverseSlices = [][]int64{
	{0, 1, 2, 3, 4, 5},
	{5, 4, 3, 2, 1, 0},
	{1, 0, 1, 1, 1, 1},
	{1, 10, 100, 1000, 10000},
}

func TestReverse(t *testing.T) {
	var expectedResults = [][]int64{
		{5, 4, 3, 2, 1, 0},
		{0, 1, 2, 3, 4, 5},
		{1, 1, 1, 1, 0, 1},
		{10000, 1000, 100, 10, 1},
	}

	for k, v := range testReverseSlices {
		t.Run(fmt.Sprintf("%d:%v", k, v), func(t *testing.T) {
			got := Reverse(testReverseSlices[k])
			if assert.ElementsMatch(t, got, testReverseSlices[k]) == false {
				t.Errorf("got: %v expected: %v given: %v", got, expectedResults[k], v)
			}
		})
	}
}
