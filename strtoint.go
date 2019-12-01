package strtoint

import (
	"fmt"
	"regexp"
	"strconv"
)

func MyStrToInt(s string) (result int, err error) {
	filteredString := Filter(s)
	result, err = strconv.Atoi(filteredString)
	return
}

func MyStrToInt2(s string) (result int, err error) {
	filteredString := Filter(s)
	_, err = fmt.Sscanf(filteredString, "%d", &result)
	return
}

func Filter(s string) (result string) {
	result = regexp.MustCompile("[^0-9]+").ReplaceAllString(s, "")
	if IsRealNegative(s) {
		result = "-" + result
	}
	return
}

func IsRealNegative(s string) (isRealNegative bool) {
	isRealNegative = false

	minusIndexes := regexp.MustCompile("-").FindStringIndex(s)
	numberIndexes := regexp.MustCompile("[0-9]").FindStringIndex(s)

	if (len(minusIndexes) > 0 && len(numberIndexes) > 0) && (minusIndexes[0] < numberIndexes[0]) {
		isRealNegative = true
	}

	return
}
