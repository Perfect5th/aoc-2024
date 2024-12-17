package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/grids"
	"github.com/Perfect5th/aoc-2024/input"
)

type robot struct {
	pos, vel grids.Point
}

// Ensures a < b && a >= 0 by 'wrapping' it within b
func clamp(a, b int) int {
	if a > b - 1 {
		return a - b
	}
	if a < 0 {
		return b + a
	}
	return a
}

// Updates `r`'s location by moving it by its velocity.
// Wraps at X w and Y h.
func (r *robot) move(w, h int) {
	newX := clamp(r.pos.X + r.vel.X, w)
	newY := clamp(r.pos.Y + r.vel.Y, h)
	r.pos = grids.NewPoint(newX, newY)
}

// Returns which of the 4 numbered quadrants `r` is in.
// quadrants are numbered clockwise, starting at 0.
// -1 means no quadrant.
func (r *robot) quadrant(w, h int) int {
	wDiv := w / 2
	hDiv := h / 2

	switch {
	case r.pos.X < wDiv && r.pos.Y < hDiv:
		return 0
	case r.pos.X > wDiv && r.pos.Y < hDiv:
		return 1
	case r.pos.X > wDiv && r.pos.Y > hDiv:
		return 2
	case r.pos.X < wDiv && r.pos.Y > hDiv:
		return 3
	}
	return -1
}

// Mutates `robots` by moving each robot, wrapping at `w` width and `h` height.
func tick(robots []*robot, w, h int) {
	for _, r := range robots {
		r.move(w, h)
	}
}

// Counts the number of `robots` in each quadrant of the `w` by `h` space.
func safetyFactor(robots []*robot, w, h int) int {
	quads := []int{0, 0, 0, 0}

	for _, r := range robots {
		quad := r.quadrant(w, h)
		if quad >= 0 {
			quads[quad]++
		}
	}

	total := 1
	for _, q := range quads {
		total *= q
	}
	return total
}

// Calculates the 'safety factor' of the robots after 100 seconds in a space
// `w` by `h` tiles in size.
func part1(robots []*robot, w, h int) int {
	for i := 0; i < 100; i++ {
		tick(robots, w, h)
	}

	return safetyFactor(robots, w, h)
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	robots := make([]*robot, 0)
	for line := range lines {
		parts := strings.Split(line, " ")
		pos := parts[0][2:]
		posXY := strings.Split(pos, ",")
		posX, err := strconv.Atoi(posXY[0])
		if err != nil {
			panic(err)
		}
		posY, err := strconv.Atoi(posXY[1])
		if err != nil {
			panic(err)
		}
		posP := grids.NewPoint(posX, posY)
		vel := parts[1][2:]
		velXY := strings.Split(vel, ",")
		velX, err := strconv.Atoi(velXY[0])
		if err != nil {
			panic(err)
		}
		velY, err := strconv.Atoi(velXY[1])
		if err != nil {
			panic(err)
		}
		velP := grids.NewPoint(velX, velY)
		r := &robot{posP, velP}
		robots = append(robots, r)
	}

	fmt.Println(part1(robots, 101, 103))
}
