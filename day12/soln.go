package main

import (
	"fmt"

	"github.com/Perfect5th/aoc-2024/grids"
	"github.com/Perfect5th/aoc-2024/input"
)

// Calculates the perimeter of `grid`. That is, the number of sides of points that
// don't have neighbours.
func perimeter[T comparable](grid *grids.Grid[T]) (p int) {
	for point := range grid.Iter() {
		_, ok := grid.Get(point)
		if !ok {
			continue
		}

		neighbours := grid.Directions(point)
		p += 4 - len(neighbours)
	}
	return
}

// This is colouring DFS - starting from `start`, traverse `grid` to find every connected point
// with the same rune.
func regionDfs(grid *grids.Grid[rune], start grids.Point, tags map[grids.Point]bool) (region *grids.Grid[bool]) {
	queue := []grids.Point{start}
	region = grids.NewGrid[bool]()
	regionType, ok := grid.Get(start)
	if !ok {
		panic("Need something to colour the region")
	}

	for len(queue) > 0 {
		curr := queue[0]

		for _, p := range grid.Directions(curr) {
			if tags[p] {
				continue
			}
			value, ok := grid.Get(p)
			if !ok || value != regionType {
				continue
			}
			region.AddP(p, true)
			tags[p] = true
			queue = append(queue, p)
		}

		queue = queue[1:]
	}

	return
}

// Assign every point in `grid` to a region made up of adjacent identical values.
func regions(grid *grids.Grid[rune]) (rs []*grids.Grid[bool]) {
	rs = make([]*grids.Grid[bool], 0)
	tags := make(map[grids.Point]bool)

	for p := range grid.Iter() {
		if tags[p] {
			continue
		}
		region := regionDfs(grid, p, tags)
		region.AddP(p, true)
		rs = append(rs, region)
	}

	return
}

// Sums the total of fencing cost for all regions of the map.
func part1(grid *grids.Grid[rune]) (total int) {
	for _, region := range regions(grid) {
		size := region.Size()
		peri := perimeter(region)
		cost := size * peri
		fmt.Printf("REGION: %d * %d = %d\n", size, peri, cost)
		total += cost
	}
	return
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	grid := grids.NewGrid[rune]()
	y := 0
	for line := range lines {
		for x, r := range []rune(line) {
			grid.AddPoint(x, y, r)
		}
		y++
	}

	fmt.Println(part1(grid))
}
