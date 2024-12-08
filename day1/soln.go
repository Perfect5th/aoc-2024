package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/utils"
)

func part1(a []int, b []int) (sum int) {
	slices.Sort(a)
	slices.Sort(b)

	for i, num := range a {
		dist := num - b[i]

		if dist < 0 {
			dist = dist * -1
		}

		sum += dist
	}
	return
}

// a and b are assumed to be sorted.
func part2(a []int, b []int) (sum int) {
	for _, num := range a {
		times := 0

		for _, other := range b {
			if num == other {
				times++
			}
		}

		sum += num * times
	}

	return
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		panic(err)
	}

	a := make([]int, 0)
	b := make([]int, 0)

	for line := range lines {
		// input is lines of pairs of integers.
		split := strings.Split(line, "   ")

		first, erra := strconv.Atoi(split[0])
		second, errb := strconv.Atoi(split[1])

		if erra != nil {
			panic(erra)
		}

		if errb != nil {
			panic(errb)
		}

		a = append(a, first)
		b = append(b, second)
	}

	fmt.Println(part1(a, b))
	fmt.Println(part2(a, b))
}
