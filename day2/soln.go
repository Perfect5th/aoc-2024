package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/utils"
)

// Returns `true` if levels are safe:
// ((all increasing) OR (all decreasing)) AND
// all adjacents differ by >= 1 AND < 3
func safe(report []int) bool {
	increasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		prev := report[i-1]
		curr := report[i]

		if curr < prev && increasing {
			return false
		}

		if curr > prev && !increasing {
			return false
		}

		diff := curr - prev
		if diff < 0 {
			diff = diff * -1
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

// Generates a slice of reports from `report` where each item is removed
func dampens(report []int) (chan []int) {
	variants := make(chan []int)

	go func() {
		for i := 0; i < len(report); i++ {
			variant := slices.Clone(report)
			slices.Delete(variant, i, i+1)
			variants <- variant[:len(report)-1]
		}

		close(variants)
	}()

	return variants
}

func part1(reports [][]int) (result int) {
	for _, report := range reports {
		if safe(report) {
			result++
		}
	}

	return
}

func part2(reports [][]int) (result int) {
	for _, report := range reports {
		if safe(report) {
			result++
		} else {
			for dampened := range dampens(report) {
				if safe(dampened) {
					result++
					break
				}
			}
		}
	}

	return
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		panic(err)
	}

	reports := make([][]int, 0)
	for line := range lines {
		values := strings.Split(line, " ")
		ints := make([]int, len(values))

		for i, v := range values {
			int, err := strconv.Atoi(v)

			if err != nil {
				panic(err)
			}

			ints[i] = int
		}

		reports = append(reports, ints)
	}

	fmt.Println(part1(reports))
	fmt.Println(part2(reports))
}
