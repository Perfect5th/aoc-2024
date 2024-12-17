package main

import (
	"fmt"
	"slices"

	"github.com/Perfect5th/aoc-2024/algos"
	"github.com/Perfect5th/aoc-2024/input"
	"github.com/Perfect5th/aoc-2024/grids"
)

type pointdir struct {
	point grids.Point
	dir int
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	var start grids.Point
	var end grids.Point

	grid := grids.NewGrid[rune]()
	y := 0
	for line := range lines {
		for x, r := range []rune(line) {
			if r == 'S' {
				start = grids.NewPoint(x, y)
			}
			if r == 'E' {
				end = grids.NewPoint(x, y)
			}
			grid.AddPoint(x, y, r)
		}
		y++
	}

	fmt.Println(part1(grid, start, end))
}

// Converts `grid` into a graph, then performs Dijkstra's shortest path algo on it.
func part1(grid *grids.Grid[rune], start grids.Point, end grids.Point) int {
	vertices, edges, source, target := traverse(grid, pointdir{start, grids.EAST}, end)
	graph := algos.NewGraph(vertices, edges)

	return int(algos.DijkstraCost(graph, source, target))
}

// Converts `grid` into a graph.
func traverse(grid *grids.Grid[rune], start pointdir, end grids.Point) (
	[]*algos.Vertex[pointdir],
	map[algos.Edge[pointdir]]float64,
	*algos.Vertex[pointdir],
	*algos.Vertex[pointdir],
) {
	var target *algos.Vertex[pointdir]
	startv := algos.NewVertex(start)

	edges := make(map[algos.Edge[pointdir]]float64)
	vertices := []*algos.Vertex[pointdir]{startv}

	queue := []pointdir{start}
	done := map[pointdir]*algos.Vertex[pointdir]{start: startv}
	// Build up all the edges first.
	for len(queue) > 0 {
		curr := queue[0]
		queue = slices.Delete(queue, 0, 1)

		currv := done[curr]

		for e, cost := range adjacents(grid, curr) {
			vertex, isDone := done[e]
			if !isDone {
				vertex = algos.NewVertex(e)
				vertices = append(vertices, vertex)
				done[e] = vertex
				queue = append(queue, e)

				if e.point == end {
					target = vertex
				}
			}
			edges[algos.Edge[pointdir]{currv, vertex}] = cost
		}
	}

	// Add neighbors
	for k, _ := range edges {
		k.U.AddNeighbor(k.V)
	}

	return vertices, edges, startv, target
}

// Produces the adjacent pointdir vertices for pointdir `v`.
func adjacents(grid *grids.Grid[rune], v pointdir) map[pointdir]float64 {
	adj := make(map[pointdir]float64)

	val, _ := grid.Get(v.point)
	if val == 'E' {
		return adj
	}

	var next grids.Point
	switch {
	case v.dir == grids.NORTH:
		adj[pointdir{v.point, grids.EAST}] = 1000.0
		adj[pointdir{v.point, grids.WEST}] = 1000.0
		next = grids.Point{v.point.X, v.point.Y-1}
	case v.dir == grids.EAST:
		adj[pointdir{v.point, grids.SOUTH}] = 1000.0
		adj[pointdir{v.point, grids.NORTH}] = 1000.0
		next = grids.Point{v.point.X+1, v.point.Y}
	case v.dir == grids.SOUTH:
		adj[pointdir{v.point, grids.WEST}] = 1000.0
		adj[pointdir{v.point, grids.EAST}] = 1000.0
		next = grids.Point{v.point.X, v.point.Y+1}
	case v.dir == grids.WEST:
		adj[pointdir{v.point, grids.NORTH}] = 1000.0
		adj[pointdir{v.point, grids.SOUTH}] = 1000.0
		next = grids.Point{v.point.X-1, v.point.Y}
	}

	val, ok := grid.Get(next)
	if ok && val != '#' {
		adj[pointdir{next, v.dir}] = 1.0
	}

	return adj
}
