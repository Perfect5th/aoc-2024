package main

import (
	"fmt"
	// "os"

	"github.com/Perfect5th/aoc-2024/grids"
	"github.com/Perfect5th/aoc-2024/input"
)

type Box struct {
	grid *grids.Grid[rune]
	l grids.Point
	r grids.Point
}

func NewBox(grid *grids.Grid[rune], position grids.Point) *Box {
	return &Box{
		grid: grid,
		l: position,
		r: grids.NewPoint(position.X+1, position.Y),
	}
}

// Returns True if box can be moved in direction `d`
func (b *Box) canPushY(d grids.Point) bool {
	// This looks a lot like `pushY`, but doesn't do the actually move.
	nextL := grids.NewPoint(b.l.X, b.l.Y+d.Y)
	nextR := grids.NewPoint(b.r.X, b.r.Y+d.Y)
	valueL, _ := b.grid.Get(nextL)
	valueR, _ := b.grid.Get(nextR)

	switch {
	case valueL == ']' && valueR == '[':
		// Two boxes to push - need to be able to push both or neither moves
		boxL := NewBox(b.grid, grids.NewPoint(nextL.X-1, nextL.Y))
		boxR := NewBox(b.grid, nextR)
		return boxL.canPushY(d) && boxR.canPushY(d)
	case valueL == '[' && valueR == ']':
		box := NewBox(b.grid, nextL)
		return box.canPushY(d)
	case valueL == ']' && valueR == '.':
		box := NewBox(b.grid, grids.NewPoint(nextL.X-1, nextL.Y))
		return box.canPushY(d)
	case valueL == '.' && valueR == '[':
		box := NewBox(b.grid, nextR)
		return box.canPushY(d)
	case valueL == '.' && valueR == '.':
		return true
	case valueL == '#' || valueR == '#':
		return false
	}

	return false
}

// Moves box in the X plane, propagating movement to other boxes.
func (b *Box) pushX(d grids.Point) bool {
	var next grids.Point
	if d.X < 0 {
		next = grids.NewPoint(b.l.X-1, b.l.Y)
	} else {
		next = grids.NewPoint(b.r.X+1, b.r.Y)
	}
	value, _ := b.grid.Get(next)
	pushed := false

	switch {
	case value == '[':
		box := NewBox(b.grid, next)
		pushed = box.push(d)
	case value == ']':
		box := NewBox(b.grid, grids.NewPoint(next.X-1, next.Y))
		pushed = box.push(d)
	case value == '.':
		pushed = true
	case value == '#':
		pushed = false
	}

	if pushed {
		newL := grids.NewPoint(b.l.X+d.X, b.l.Y)
		newR := grids.NewPoint(b.r.X+d.X, b.r.Y)
		b.grid.Set(b.l, '.')
		b.grid.Set(b.r, '.')
		b.grid.Set(newL, '[')
		b.grid.Set(newR, ']')
	}

	return pushed
}

// Moves box in the Y plane, propagating movement to other boxes.
func (b *Box) pushY(d grids.Point) bool {
	nextL := grids.NewPoint(b.l.X, b.l.Y+d.Y)
	nextR := grids.NewPoint(b.r.X, b.r.Y+d.Y)
	valueL, _ := b.grid.Get(nextL)
	valueR, _ := b.grid.Get(nextR)
	pushed := false

	switch {
	case valueL == ']' && valueR == '[':
		// Two boxes to push - need to be able to push both or neither moves
		boxL := NewBox(b.grid, grids.NewPoint(nextL.X-1, nextL.Y))
		boxR := NewBox(b.grid, nextR)
		if boxL.canPushY(d) && boxR.canPushY(d) {
			boxL.push(d)
			boxR.push(d)
			pushed = true
		}
	case valueL == '[' && valueR == ']':
		box := NewBox(b.grid, nextL)
		pushed = box.push(d)
	case valueL == ']' && valueR == '.':
		box := NewBox(b.grid, grids.NewPoint(nextL.X-1, nextL.Y))
		pushed = box.push(d)
	case valueL == '.' && valueR == '[':
		box := NewBox(b.grid, nextR)
		pushed = box.push(d)
	case valueL == '.' && valueR == '.':
		pushed = true
	case valueL == '#' || valueR == '#':
		pushed = false
	}

	if pushed {
		newL := grids.NewPoint(b.l.X, b.l.Y+d.Y)
		newR := grids.NewPoint(b.r.X, b.r.Y+d.Y)
		b.grid.Set(b.l, '.')
		b.grid.Set(b.r, '.')
		b.grid.Set(newL, '[')
		b.grid.Set(newR, ']')
	}

	return pushed
}

// Moves box in direction specified by `d`. If it can't be moved,
// returns false. This propagates movement to boxes above this one.
func (b *Box) push(d grids.Point) bool {
	if d.X != 0 {
		return b.pushX(d)
	}
	if d.Y != 0 {
		return b.pushY(d)
	}
	return false
}

type Robot struct {
	grid *grids.Grid[rune]
	moves []rune
	position grids.Point
	moveCount int
}

func NewRobot(grid *grids.Grid[rune], moves []rune, start grids.Point) *Robot {
	return &Robot{grid: grid, moves: moves, position: start}
}

// Move in the direction specified by `d`, propagating the motion ot any connected boxes.
func (r *Robot) pushAndMove(d grids.Point) {
	moveTo := grids.NewPoint(r.position.X + d.X, r.position.Y + d.Y)
	next := moveTo
	value, _ := r.grid.Get(next)
	pushing := false

	// Search for boxes
	for value != '#' && value != '.' {
		pushing = true
		next = grids.NewPoint(next.X + d.X, next.Y + d.Y)
		value, _ = r.grid.Get(next)
	}

	// Can't move
	if value == '#' {
		return
	}

	// Value is '.'
	if pushing {
		r.grid.Set(next, 'O')
		r.grid.Set(moveTo, '.')
	}

	r.position = moveTo
}

func (r *Robot) pushAndMove2(d grids.Point) {
	next := grids.NewPoint(r.position.X + d.X, r.position.Y + d.Y)
	value, _ := r.grid.Get(next)
	moved := false

	switch {
	case value == '[':
		box := NewBox(r.grid, next)
		moved = box.push(d)
	case value == ']':
		box := NewBox(r.grid, grids.NewPoint(next.X-1, next.Y))
		moved = box.push(d)
	case value == '.':
		moved = true
	}

	if moved {
		r.position = next
	}
}

// Mutates `r`'s state by going through its move list until complete,
// pushing boxes, etc.
func (r *Robot) move(moveFunc func(grids.Point)) bool {
	if r.moveCount < len(r.moves) {
		var delta grids.Point
		m := r.moves[r.moveCount]

		switch {
		case m == '^':
			delta = grids.NewPoint(0, -1)
		case m == '>':
			delta = grids.NewPoint(1, 0)
		case m == 'v':
			delta = grids.NewPoint(0, 1)
		default:
			delta = grids.NewPoint(-1, 0)
		}
		moveFunc(delta)

		r.moveCount++
		return true
	}

	return false
}

// Calculates the Goods Position System value for the given `point`.
func gps(point grids.Point) int {
	return (100 * point.Y) + point.X
}

// Returns the sum of the GPS coordinates of the boxes after the robot
// has finished moving.
func part1(grid *grids.Grid[rune], robot *Robot) (sum int) {
	for robot.move(robot.pushAndMove) {}

	for point := range grid.Iter() {
		value, _ := grid.Get(point)
		if value == 'O' {
			sum += gps(point)
		}
	}
	return
}

// Returns the sum of the GPS coordinates of the boxes after the robot
// has finished moving.
func part2(grid *grids.Grid[rune], robot *Robot) (sum int) {
	// i := 0
	// buff := make([]byte, 1024)
	// for _, line := range grid.Lines() {
	// 	fmt.Println(string(line))
	// }
	// for robot.move(robot.pushAndMove2) {
	// 	fmt.Println(i)
	// 	fmt.Println(string(robot.moves[i]))
	// 	i++
	// 	fmt.Println(robot.position)
	// 	for y, line := range grid.Lines() {
	// 		if y == robot.position.Y {
	// 			line[robot.position.X] = '@'
	// 		}
	// 		fmt.Println(string(line))
	// 	}
	// 	_, _ = os.Stdin.Read(buff)
	// }
	for robot.move(robot.pushAndMove2) {}

	for point := range grid.Iter() {
		value, _ := grid.Get(point)
		if value == '[' {
			sum += gps(point)
		}
	}
	return
}

func main() {
	lines, err := input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	var start grids.Point
	grid := grids.NewGrid[rune]()
	y := 0
	for line := range lines {
		if line == "" {
			break
		}
		for x, r := range []rune(line) {
			if r == '@' {
				start = grids.NewPoint(x, y)
				r = '.'
			}
			grid.AddPoint(x, y, r)
		}
		y++
	}

	moves := make([]rune, 0)
	for line := range lines {
		moves = append(moves, []rune(line)...)
	}

	robot := NewRobot(grid, moves, start)
	fmt.Println(part1(grid, robot))

	lines, err = input.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	var start2 grids.Point
	grid2 := grids.NewGrid[rune]()
	y = 0
	for line := range lines {
		if line == "" {
			break
		}
		x := 0
		for _, r := range []rune(line) {
			switch {
			case r == '@':
				start2 = grids.NewPoint(x, y)
				grid2.AddPoint(x, y, '.')
				grid2.AddPoint(x+1, y, '.')
			case r == 'O':
				grid2.AddPoint(x, y, '[')
				grid2.AddPoint(x+1, y, ']')
			case r == '#':
				grid2.AddPoint(x, y, '#')
				grid2.AddPoint(x+1, y, '#')
			case r == '.':
				grid2.AddPoint(x, y, '.')
				grid2.AddPoint(x+1, y, '.')
			}
			x += 2
		}
		y++
	}

	moves2 := make([]rune, 0)
	for line := range lines {
		moves2 = append(moves2, []rune(line)...)
	}

	robot2 := NewRobot(grid2, moves2, start2)
	fmt.Println(part2(grid2, robot2))
}
