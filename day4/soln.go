package main

import (
	"fmt"

	"github.com/Perfect5th/aoc-2024/utils"
)

// NOTE: the 'tester' functions below assume bounds have been checked, and that 'M'
// already appears.
// Returns `true` if "XMAS" appears in `grid` from left to right.
func rightXmas(x, y int, grid [][]rune) bool {
	return grid[y][x+2] == 'A' && grid[y][x+3] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from top to bottom.
func downXmas(x, y int, grid [][]rune) bool {
	return grid[y+2][x] == 'A' && grid[y+3][x] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from right to left.
func leftXmas(x, y int, grid [][]rune) bool {
		return grid[y][x-2] == 'A' && grid[y][x-3] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from bottom to top.
func upXmas(x, y int, grid [][]rune) bool {
		return grid[y-2][x] == 'A' && grid[y-3][x] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from bottom left to top right.
func rightUpXmas(x, y int, grid [][]rune) bool {
	return grid[y-2][x+2] == 'A' && grid[y-3][x+3] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from top left to bottom right.
func rightDownXmas(x, y int, grid [][]rune) bool {
	return grid[y+2][x+2] == 'A' && grid[y+3][x+3] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from top right to bottom left.
func leftDownXmas(x, y int, grid [][]rune) bool {
	return grid[y+2][x-2] == 'A' && grid[y+3][x-3] == 'S'
}

// Returns `true` if "XMAS" appears in `grid` from bottom right to top left.
func leftUpXmas(x, y int, grid [][]rune) bool {
	return grid[y-2][x-2] == 'A' && grid[y-3][x-3] == 'S'
}

const (
	UP = iota
	RIGHTUP
	RIGHT
	RIGHTDOWN
	DOWN
	LEFTDOWN
	LEFT
	LEFTUP
)

var testers = map[int]func(int, int, [][]rune) bool{
	UP:        upXmas,
	RIGHTUP:   rightUpXmas,
	RIGHT:     rightXmas,
	RIGHTDOWN: rightDownXmas,
	DOWN:      downXmas,
	LEFTDOWN:  leftDownXmas,
	LEFT:      leftXmas,
	LEFTUP:    leftUpXmas,
}

// Returns a list of directions from `x`,`y` in `grid` that could contain `"MAS"`.
func findMs(x, y int, grid [][]rune) (directions []int) {
	if y > 2 && grid[y-1][x] == 'M' {
		directions = append(directions, UP)
	}
	if x < len(grid[0]) - 3 && y > 2 && grid[y-1][x+1] == 'M' {
		directions = append(directions, RIGHTUP)
	}
	if x < len(grid[0]) - 3 && grid[y][x+1] == 'M' {
		directions = append(directions, RIGHT)
	}
	if x < len(grid[0]) - 3 && y < len(grid) - 3 && grid[y+1][x+1] == 'M' {
		directions = append(directions, RIGHTDOWN)
	}
	if y < len(grid) - 3 && grid[y+1][x] == 'M' {
		directions = append(directions, DOWN)
	}
	if x > 2 && y < len(grid) - 3 && grid[y+1][x-1] == 'M' {
		directions = append(directions, LEFTDOWN)
	}
	if x > 2 && grid[y][x-1] == 'M' {
		directions = append(directions, LEFT)
	}
	if x > 2 && y > 2 && grid[y-1][x-1] == 'M' {
		directions = append(directions, LEFTUP)
	}
	return
}

// Returns the number of `"XMAS"`es starting at grid location `x`,`y`, going in any direction.
func xmases(x, y int, grid [][]rune) (count int) {
	directions := findMs(x, y, grid)
	if len(directions) == 0 {
		return
	}

	for _, direction := range directions {
		if testers[direction](x, y, grid) {
			count++
		}
	}
	return
}

// Returns the number of times "XMAS" appears in `grid`, in any direction.
func part1(grid [][]rune) (count int) {
	for y, row := range grid {
		for x, char := range row {
			if char == 'X' {
				count += xmases(x, y, grid)
			}
		}
	}
	return
}

// Returns 1 if there is a "MAS", X-crossed on the 'A', centered at `x,y` in `grid`.
// Otherwise returns 0.
func crossMases(x, y int, grid [][]rune) (xmas bool) {
	if x < 1 || y < 1 {
		return
	}

	if x > len(grid[0]) - 2 || y > len(grid) - 2 {
		return
	}

	upLeft := grid[y-1][x-1]
	downRight := grid[y+1][x+1]
	upRight := grid[y-1][x+1]
	downLeft := grid[y+1][x-1]

	if (upLeft == 'M' && downRight == 'S') || (upLeft == 'S' && downRight == 'M') {
		xmas = (upRight == 'M' && downLeft == 'S') || (upRight == 'S' && downLeft == 'M')
	}

	return
}

// Returns the number of times "MAS" appears, X-crossed on the 'A', in `grid`
func part2(grid [][]rune) (count int) {
	for y, row := range grid {
		for x, char := range row {
			if char == 'A' && crossMases(x, y, grid) {
				count++
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

	grid := make([][]rune, 0)
	for line := range lines {
		grid = append(grid, []rune(line))
	}

	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}
