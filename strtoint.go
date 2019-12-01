package strtoint

import (
	"fmt"
	"regexp"
	"strconv"
)

func MyStrToInt(s string) (result int, err error) {
	result, err = strconv.Atoi(Filter(s))
	return
}

func MyStrToInt2(s string) (result int, err error) {
	_, err = fmt.Sscanf(Filter(s), "%d", &result)
	return
}

func Filter(s string) (result string) {
	result = regexp.MustCompile("[^0-9]+").ReplaceAllString(s, "")
	if IsRealNegative(s) {
		result = "-" + result
	}
	return
}

func IsRealNegative(s string) bool {
	minusIndexes := regexp.MustCompile("-").FindStringIndex(s)
	numberIndexes := regexp.MustCompile("[0-9]").FindStringIndex(s)

	if (len(minusIndexes) > 0 && len(numberIndexes) > 0) && (minusIndexes[0] < numberIndexes[0]) {
		return true
	}

	return false
}
