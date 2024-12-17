package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/input"
)

// Returns the left half and right half of `value`, as integers.
// It is assumed that `value` has an even length.
func split(value string) (left, right int) {
	var err error
	leftStr := value[len(value)/2:]
	rightStr := value[:len(value)/2]
	left, err = strconv.Atoi(leftStr)
	if err != nil {
		panic(err)
	}
	right, err = strconv.Atoi(rightStr)
	if err != nil {
		panic(err)
	}
	return
}

// Mutates `l` (a list of ints) by operting on each item, following these rules:
// - if 0, change to 1
// - if even number of digits, change to last half of digits, stripping leading zeroes
//     insert in front a new element with first half of digits
// - otherwise, multiply value by 2024
func blink(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		value := e.Value.(int)
		if value == 0 {
			e.Value = 1
			continue
		}

		valueStr := fmt.Sprintf("%d", value)
		if len(valueStr) % 2 == 0 {
			left, right := split(valueStr)
			l.InsertBefore(left, e)
			e.Value = right
			continue
		}

		e.Value = e.Value.(int) * 2024
	}
}

// Returns the length of `l` after "blinking" 25 times.
func part1(l *list.List) int {
	for i := 0; i < 25; i++ {
		blink(l)
	}
	return l.Len()
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	l := list.New()
	// There's only one line again.
	for line := range lines {
		parts := strings.Split(line, " ")
		for _, p := range parts {
			pi, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			l.PushBack(pi)
		}
	}

	fmt.Println(part1(l))
}
