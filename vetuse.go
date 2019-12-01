package strtoint

import (
	"encoding/json"
	"errors"
	"fmt"
)

func multiplyByTwo(k int) (err error) {
	k *= 2
	return nil
}

func printMoreTen(k int64) (err error) {
	if k < 10 {
		return errors.New("k must be > 10")
	}
	fmt.Println(k)
	return
}

func dejson() (err error) {
	type jsStruct struct {
		Data int  `json:"data"`
		OK   bool `json:"ok"`
	}
	in := []byte(`{"data": 13, "ok": true}`)
	var out jsStruct
	if err := json.Unmarshal(in, &out); err != nil {
		panic(err)
	}

	return err
}

func main() {
	var r int64 = 11
	if err := printMoreTen(r); err != nil {
		panic(err)
	}
}
