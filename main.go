package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].birthDay == p[j].birthDay {
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.Before(p[j].birthDay)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) String() (result string) {
	for _, v := range p {
		result += fmt.Sprintf("%v | %s %s\n", v.birthDay, v.firstName, v.lastName)
	}

	result += "\n"
	return
}

func (p People) Sort() {
	sort.Sort(p)
}
