package main

import (
	"fmt"
	"sort"
)

func Average(inputArray []int) float64 {
	return float64(CalculateSumOfArray(inputArray)) / float64(len(inputArray))
}

func CalculateSumOfArray(inputArray []int) (result int) {
	for _, v := range inputArray {
		result += v
	}
	return
}

func Max(inputSlice []string) (maxString string) {
	if len(inputSlice) == 0 {
		return ""
	}

	for _, v := range inputSlice {
		if len(v) > len(maxString) {
			maxString = v
		}
	}
	return
}

func Reverse(inputSlice []int64) []int64 {
	var reversedSlice = make([]int64, 0, cap(inputSlice))
	for i := len(inputSlice) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, inputSlice[i])
	}

	return reversedSlice
}

/*
fmt.Printf("%v", inputMap) - https://tip.golang.org/doc/go1.12#fmt
*/
func PrintSorted(inputMap map[int]string) {
	keys := make([]int, 0, len(inputMap))
	for k := range inputMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Printf("%s", inputMap[v])
	}
}

func main() {
	var sortMaps = []map[int]string{
		{2: "b", 3: "c", 1: "a"},
		{1000: "c", 100: "b", 1: "a"},
	}

	for _, v := range sortMaps {
		PrintSorted(v)
	}
}
