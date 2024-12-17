package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Perfect5th/aoc-2024/input"
	"github.com/Perfect5th/aoc-2024/grids"
)

type machine struct {
	a, b, prize grids.Point
}

// Parses `desc` into a point describing the behaviour of the button.
func newButton(desc string) grids.Point {
	parts := strings.Split(desc, " ")
	xStr := parts[2]
	xStr = xStr[2:len(xStr)-1]
	yStr := parts[3]
	yStr = yStr[2:]

	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}

	return grids.NewPoint(x, y)
}

// Parses `desc` into a point describing the location of a prize.
func newPrize(desc string) grids.Point {
	parts := strings.Split(desc, " ")
	xStr := parts[1]
	xStr = xStr[2:len(xStr)-1]
	yStr := parts[2]
	yStr = yStr[2:]

	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}

	return grids.NewPoint(x, y)
}

func newMachine(a string, b string, p string) *machine {
	return &machine{
		a: newButton(a),
		b: newButton(b),
		prize: newPrize(p),
	}
}

// Finds the lowest cost way to win `m`'s prize. Returns `_, false` if unwinnable.
// This requires solving a set of two linear equations, minimizing the value for A.
// We do this using the elimination method.
func (m *machine) win() (int, bool) {
	xPrizeMult := m.prize.X * m.a.Y
	yPrizeMult := m.prize.Y * m.a.X
	xBMult := m.b.X * m.a.Y
	yBMult := m.b.Y * m.a.X

	bDiff := xBMult - yBMult
	prizeDiff := xPrizeMult - yPrizeMult

	// We need to detect if the system is unsolveable.
	if prizeDiff % bDiff != 0 {
		return 0, false
	}

	bPresses := prizeDiff / bDiff
	aNum := m.prize.X - (m.b.X * bPresses)

	if aNum % m.a.X != 0 {
		return 0, false
	}
	aPresses := aNum / m.a.X

	return (aPresses * 3) + bPresses, true
}

// Finds the lowest cost to win all prizes in all `machines`
func part1(machines []*machine) (cost int) {
	for _, machine := range machines {
		c, won := machine.win()
		if won {
			cost += c
		}
	}
	return
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	machines := make([]*machine, 0)
	for a := range lines {
		b := <-lines
		p := <-lines
		<-lines

		machines = append(machines, newMachine(a, b, p))
	}

	fmt.Println(part1(machines))
}
