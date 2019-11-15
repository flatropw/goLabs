package main

import (
	"fmt"
)

func main() {
	/*
		pay attention: i = -1 will throw overflow error
	*/
	factorialTestSet := [...]uint{0, 1, 5, 6, 65, 66}
	for _, v := range factorialTestSet {
		fmt.Println(fmt.Sprintf("%d! = %d", v, calculateFactorial(v)))
	}
}

func calculateFactorial(i uint) (result uint) {
	result = 1
	for counter := uint(1); counter <= i; counter++ {
		result *= counter
	}
	return
}
