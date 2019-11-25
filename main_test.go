package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestPeople_Sort(t *testing.T) {
	IvanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-09")
	AlexAlexovDate, err := time.Parse("2006-Jan-02", "2004-Aug-09")
	SergeySergeevDate, err := time.Parse("2006-Jan-02", "2004-Aug-09")
	JamesBondDate, err := time.Parse("2006-Jan-02", "2011-Aug-09")
	ThamesBondDate, err := time.Parse("2006-Jan-02", "2011-Aug-09")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	given := People{
		{"BIvan", "Ivanov", IvanIvanovDate},
		{"Alex", "Alexov", AlexAlexovDate},
		{"CSergey", "Sergeev", SergeySergeevDate},
		{"James", "Bond", JamesBondDate},
		{"Thames", "Bond", ThamesBondDate},
	}

	expected := People{
		{"Alex", "Alexov", AlexAlexovDate},
		{"CSergey", "Sergeev", SergeySergeevDate},
		{"BIvan", "Ivanov", IvanIvanovDate},
		{"James", "Bond", JamesBondDate},
		{"Thames", "Bond", ThamesBondDate},
	}

	t.Run("People Sort by bday fname", func(t *testing.T) {
		got := given
		sort.Sort(got)
		if reflect.DeepEqual(got, expected) == false {
			t.Errorf("got: %v expected: %v given: %v", got, expected, given)
		}
	})
}
