package main

import (
	"fmt"
	"maps"

	"github.com/Perfect5th/aoc-2024/utils"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type Point struct {x, y int}

type Guard struct {
	facing int
	starting Point
	position Point
	grid map[Point]rune
	obstacle *Point
	trail map[Point]map[int]bool
}

func NewGuard(grid map[Point]rune) Guard {
	trail := make(map[Point]map[int]bool)
	for k, v := range grid {
		if v == '^' {
			return Guard{
				starting: k,
				position: k,
				grid: grid,
				trail: trail,
			}
		}
	}

	return Guard{grid: grid, trail: trail}
}

func (g *Guard) Clone() Guard {
	guard := Guard{
		g.facing,
		g.starting,
		g.position,
		g.grid,
		nil,
		make(map[Point]map[int]bool),
	}

	for k, v := range g.trail {
		clone := maps.Clone(v)
		guard.trail[k] = clone
	}

	return guard
}

func (g *Guard) Front() Point {
	switch {
	case g.facing == NORTH:
		return Point{g.position.x, g.position.y-1}
	case g.facing == EAST:
		return Point{g.position.x+1, g.position.y}
	case g.facing == SOUTH:
		return Point{g.position.x, g.position.y+1}
	default:
		return Point{g.position.x-1, g.position.y}
	}
}

func (g *Guard) AddObstacle() {
	front := g.Front()
	_, ok := g.grid[front]

	if ok {
		g.obstacle = &front
	}
}

func (g *Guard) CanStep() bool {
	_, ok := g.grid[g.position]
	return ok
}

func (g *Guard) turn() {
	switch {
	case g.facing == NORTH:
		g.facing = EAST
	case g.facing == EAST:
		g.facing = SOUTH
	case g.facing == SOUTH:
		g.facing = WEST
	case g.facing == WEST:
		g.facing = NORTH
	}
}

func (g *Guard) step() {
	switch {
	case g.facing == NORTH:
		g.position.y--
	case g.facing == EAST:
		g.position.x++
	case g.facing == SOUTH:
		g.position.y++
	case g.facing == WEST:
		g.position.x--
	}
}

func (g *Guard) addTrail() {
	pos, ok := g.trail[g.position]

	if ok {
		pos[g.facing] = true
	} else {
		g.trail[g.position] = make(map[int]bool)
		g.trail[g.position][g.facing] = true
	}
}

func (g *Guard) Step() {
	front := g.Front()
	object := g.grid[front]

	g.addTrail()

	if object == '#' {
		g.turn()
		g.Step()
	} else {
		g.step()
	}
}

func (g *Guard) Looping() bool {
	pos, ok := g.trail[g.position]

	if !ok {
		return false
	}

	return pos[g.facing]
}

func canLoop(g *Guard) bool {
	newGuard := g.Clone()
	if newGuard.Front() == newGuard.starting {
		return false
	}
	newGuard.AddObstacle()

	for newGuard.CanStep() {
		newGuard.Step()

		if newGuard.Looping() {
			return true
		}
	}

	return false
}

func solve(grid map[Point]rune) (int, int) {
	guard := NewGuard(grid)
	obstacles := make(map[Point]bool)

	for guard.CanStep() {
		guard.Step()
	}

	for position, _ := range guard.trail {
		if position == guard.starting {
			continue
		}

		newGrid := maps.Clone(grid)
		newGrid[position] = '#'
		newGuard := NewGuard(newGrid)

		for newGuard.CanStep() {
			newGuard.Step()

			if newGuard.Looping() {
				obstacles[position] = true
				break
			}
		}
	}

	return len(guard.trail), len(obstacles)
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	grid := make(map[Point]rune, 0)
	y := 0
	for line := range lines {
		for x, rune := range []rune(line) {
			grid[Point{x, y}] = rune
		}
		y++
	}

	part1, part2 := solve(grid)
	fmt.Println(part1)
	fmt.Println(part2)
}
