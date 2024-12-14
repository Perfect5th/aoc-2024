package main

import (
	"fmt"

	"github.com/Perfect5th/aoc-2024/utils"
)

type Point struct {x, y int}

func (p *Point) getAntinode(p2 *Point, grid map[Point]rune) *Point {
	dx := p2.x - (p.x - p2.x)
	dy := p2.y - (p.y - p2.y)

	antinode := &Point{dx, dy}

	if _, ok := grid[*antinode]; ok {
		return antinode
	}
	return nil
}

// Starting from `p`, if it is an antenna, search `grid` for matching antennaes and
// enumerate all the resulting antinodes.
func (p *Point) FindAntinodes(r rune, grid map[Point]rune) (antinodes []*Point) {
	if r == '.' {
		return
	}

	for p2, r2 := range grid {
		if *p == p2 {
			continue
		}

		if r == r2 {
			antinode := p.getAntinode(&p2, grid)
			if antinode != nil {
				antinodes = append(antinodes, antinode)
			}
		}
	}
	return
}

// Counts the unique antinode locations in `grid`.
func part1(grid map[Point]rune) int {
	antinodes := make(map[Point]bool)

	for p, r := range grid {
		as := p.FindAntinodes(r, grid)
		for _, a := range as {
			antinodes[*a] = true
		}
	}

	return len(antinodes)
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	grid := make(map[Point]rune)
	y := 0
	for line := range lines {
		for x, rune := range []rune(line) {
			grid[Point{x, y}] = rune
		}
		y++
	}

	fmt.Println(part1(grid))
}
