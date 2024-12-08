package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/utils"
)

// Returns `true` if `update` is ordered correctly according to `rules`.
// `rules` is a map of each page number to the page numbers that have to
// be before it, if they are also present.
func isCorrect(update []int, rules map[int][]int) bool {
	for i, num := range update[:len(update)-1] {
		for _, num2 := range update[i+1:] {
			rule, ok := rules[num]
			if !ok {
				continue
			}

			if slices.Contains(rule, num2) {
				return false
			}
		}
	}

	return true
}

// Sums the middle page numbers of correctly-ordered `updates`.
// An update is correctly-ordered if the ordering of its page numbers conforms to the
// `rules`.
func part1(updates [][]int, rules map[int][]int) (sum int) {
	for _, update := range updates {
		if isCorrect(update, rules) {
			sum += update[len(update) / 2]
		}
	}
	return
}

// Corrects an incorrectly-ordered `update` by sorting items according `rules`.
// SIDE-EFFECT: this mutates `update` in-place
func correctUpdate(update []int, rules map[int][]int) {
	cmp := func(a, b int) int {
		if slices.Contains(rules[a], b) {
			// b should be before a
			return 1
		}
		if slices.Contains(rules[b], a) {
			// a should be before b
			return -1
		}
		return 0
	}

	slices.SortFunc(update, cmp)
}

// Sums the middle page numbers for incorrectly-ordered `updates`, after correcting them.
// See above for defn of "correctly-ordered"
func part2(updates [][]int, rules map[int][]int) (sum int) {
	for _, update := range updates {
		if !isCorrect(update, rules) {
			correctUpdate(update, rules)
			sum += update[len(update) / 2]
		}
	}
	return
}

func main() {
	updates := make([][]int, 0)
	rules := make(map[int][]int)

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	for line := range lines {
		if strings.ContainsRune(line, '|') {
			parts := strings.Split(line, "|")
			before, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			after, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			rule, ok := rules[after]
			if ok {
				rules[after] = append(rule, before)
			} else {
				rules[after] = []int{before}
			}
		} else if len(line) > 0 {
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}

				update[i] = num
			}
			updates = append(updates, update)
		}
	}

	fmt.Println(part1(updates, rules))
	fmt.Println(part2(updates, rules))
}
