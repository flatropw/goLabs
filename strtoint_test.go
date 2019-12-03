package strtoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	given    string
	expected int
}{
	{"-346", -346},
	{"250", 250},
	{"1080", 1080},
}

var badTests = []struct {
	given    string
	expected int
}{
	{"abcd", 0},
	{"x0b", 0},
	{"$%#%ß∂∂ß˚∂µ8328™ª•™", 0},
}

func TestMyStrToInt(t *testing.T) {
	for _, test := range tests {
		result, err := MyStrToInt(test.given)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestMyStrToInt2(t *testing.T) {
	for _, test := range tests {
		result, err := MyStrToInt2(test.given)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestMyStrToIntBadData(t *testing.T) {
	for _, test := range badTests {
		result, err := MyStrToInt(test.given)
		assert.Error(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestMyStrToInt2BadData(t *testing.T) {
	for _, test := range badTests {
		result, err := MyStrToInt2(test.given)
		assert.Error(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func BenchmarkMyStrToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyStrToInt("-25fs25")
	}
}

func BenchmarkMyStrToInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyStrToInt2("-25fs25")
	}
}
