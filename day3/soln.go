package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/utils"
)

var pattern = regexp.MustCompile(`^\(\d+,\d+\)`)

// Multiplies the numbers in a string with format "\d+,\d+".
func multiply(nums string) int {
	parts := strings.Split(nums, ",")

	a, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return a * b
}

// If `mul` is a valid mul command, returns its result. Else, returns 0.
// We already know `mul` starts with `"mul"`, so we only need to check the next bit.
// `length` is the length of the match, so we can advance beyond it. It's 1 for non-matches
func validMul(mul string) (result int, length int) {
	length = 1

	match := pattern.FindString(mul)
	if len(match) > 0 {
		result = multiply(match[1:len(match)-1])
		length = len(match)
	}

	return
}

// Sums the total of (valid) mul commands in `command` by scanning the string.
func sumMuls(command string) (sum int) {
	for i := 0; i + 3 < len(command); {
		if command[i:i+3] != "mul" {
			i++
			continue
		}

		value, length := validMul(command[i+3:])
		sum += value
		i += length
	}

	return
}

// Sums the total of (valid) mul commands in `command` by scanning the string.
// Ignores muls while `enabled` is `false`.
func sumMuls2(command string, enabled bool) (sum int, curEnabled bool) {
	curEnabled = enabled

	for i := 0; i + 8 < len(command); {
		if curEnabled && command[i:i+3] == "mul" {
			value, length := validMul(command[i+3:])
			sum += value
			i += length
		} else if !curEnabled && command[i:i+4] == "do()" {
			curEnabled = true
			i += 4
		} else if curEnabled && command[i:i+7] == "don't()" {
			curEnabled = false
			i += 7
		} else {
			i++
		}
	}

	return
}

// Sums the total of (valid) mul commands in `commands`.
func part1(commands []string) (sum int) {
	for _, command := range commands {
		sum += sumMuls(command)
	}

	return
}

func part2(commands []string) (sum int) {
	enabled := true
	for _, command := range commands {
		result, newEnabled := sumMuls2(command, enabled)
		sum += result
		enabled = newEnabled
	}
	return
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		panic(err)
	}

	commands := make([]string, 0)
	for line := range lines {
		commands = append(commands, line)
	}

	fmt.Println(part1(commands))
	fmt.Println(part2(commands))
}
