package main

import (
	"fmt"
	"slices"

	"github.com/Perfect5th/aoc-2024/input"
	"github.com/Perfect5th/aoc-2024/grids"
)

type bleh struct {x, y int}

// Returns a slice of possible transitions for the current path. The last item
// in `path` is the current location. We search around it for a position that is
// exactly 1 value greater than the current spot.
func possibles(grid *grids.Grid[int], path []grids.Point) (ps []grids.Point) {
	ps = make([]grids.Point, 0)
	curr := path[len(path)-1]
	height, ok := grid.Get(curr)
	if !ok {
		return
	}

	for _, d := range grid.Directions(curr) {
		if slices.Contains(path, d) {
			continue
		}
		nh, ok := grid.Get(d)
		if !ok {
			continue
		}
		if nh - height == 1 {
			ps = append(ps, d)
		}
	}

	return
}

// Returns the number of 9-height positions reachable from `head`.
// This is effectively a type of BFS or DFS.
func score(grid *grids.Grid[int], head grids.Point) int {
	queue := make([][]grids.Point, 0)
	queue = append(queue, []grids.Point{head})
	nines := make(map[grids.Point]bool)

	for len(queue) > 0 {
		path := queue[0]
		poss := possibles(grid, path)

		for _, p := range poss {
			height, ok := grid.Get(p)
			if ok && height == 9 {
				nines[p] = true
			} else if ok {
				newPath := slices.Clone(path)
				newPath = append(newPath, p)
				queue = append(queue, newPath)
			}
		}

		queue = queue[1:]
	}

	return len(nines)
}

//  Calculates the sum of the scores of all the trailheads in `grid`.
func part1(grid *grids.Grid[int]) (sum int) {
	for point := range grid.Iter() {
		val, _ := grid.Get(point)
		if val == 0 {
			sum += score(grid, point)
		}
	}
	return
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	grid := grids.NewGrid[int]()
	y := 0
	for line := range lines {
		for x, r := range []byte(line) {
			grid.AddPoint(x, y, int(r - 48))
		}
		y++
	}

	fmt.Println(part1(grid))
}
