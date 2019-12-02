package strtoint

import (
	"fmt"
	"strconv"
)

func MyStrToInt(s string) (result int, err error) {
	result, err = strconv.Atoi(s)
	return
}

func MyStrToInt2(s string) (result int, err error) {
	_, err = fmt.Sscanf(s, "%d", &result)
	return
}

