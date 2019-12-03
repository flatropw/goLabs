package strtoint

import (
	"fmt"
	"strconv"
)

func MyStrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func MyStrToInt2(s string) (result int, err error) {
	_, err = fmt.Sscanf(s, "%d", &result)
	return
}

