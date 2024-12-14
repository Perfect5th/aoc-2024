package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/utils"
)

type Equation struct {
	Test int
	rem []int
}

// Returns `true` if `eq` can be made true using operators + and *.
func (eq *Equation) CanBeTrue() bool {
	if len(eq.rem) == 0 {
		return false
	}

	if len(eq.rem) == 1 {
		return eq.Test == eq.rem[0]
	}

	if len(eq.rem) == 2 {
		return (eq.Test == eq.rem[0] + eq.rem[1] ||
			eq.Test == eq.rem[0] * eq.rem[1])
	}

	// Recur
	plusEq := Equation{eq.Test, []int{eq.rem[0] + eq.rem[1]}}
	plusEq.rem = append(plusEq.rem, eq.rem[2:]...)

	if plusEq.CanBeTrue() {
		return true
	}

	multEq := Equation{eq.Test, []int{eq.rem[0] * eq.rem[1]}}
	multEq.rem = append(multEq.rem, eq.rem[2:]...)

	if multEq.CanBeTrue() {
		return true
	}

	return false
}

// Returns the sum of the results of `equations` that can be made true by inserting operators
// + and *.
func part1(eqs []Equation) (total int) {
	for _, eq := range eqs {
		if eq.CanBeTrue() {
			total += eq.Test
		}
	}
	return
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	eqs := make([]Equation, 0)
	for line := range lines {
		parts := strings.Split(line, " ")
		testStr := parts[0][:len(parts[0])-1]
		test, err := strconv.Atoi(testStr)
		if err != nil {
			panic(err)
		}

		rem := make([]int, len(parts)-1)
		for i, rStr := range parts[1:] {
			r, err := strconv.Atoi(rStr)
			if err != nil {
				panic(err)
			}
			rem[i] = r
		}

		eqs = append(eqs, Equation{test, rem})
	}

	fmt.Println(part1(eqs))
}
