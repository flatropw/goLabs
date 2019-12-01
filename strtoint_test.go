package strtoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	given      string
	expected   int
	isNegative bool
	filtered   string
}{
	{"----346", -346, true, "-346"},
	{"100", 100, false, "100"},
	{"abcdf-55-55+12", -555512, true, "-555512"},
	{"-0", -0, true, "-0"},
	{"3-0", 30, false, "30"},
	{"01👨‍💻25", 125, false, "0125"},
	{"ћћыƒџј01👨‍💻25", 125, false, "0125"},
	{"-0-1", -1, true, "-01"},
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

func TestFilter(t *testing.T) {
	for _, test := range tests {
		result := Filter(test.given)
		assert.Equal(t, test.filtered, result)
	}
}

func TestIsRealNegative(t *testing.T) {
	for _, test := range tests {
		result := IsRealNegative(test.given)
		assert.Equal(t, test.isNegative, result)
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
