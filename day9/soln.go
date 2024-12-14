package main

import (
	"fmt"

	"github.com/Perfect5th/aoc-2024/utils"
)

// Finds the first free spot in `disk`, starting from `l`.
func findFree(disk []int, l int) int {
	for i := l; i < len(disk); i++ {
		if disk[i] == -1 {
			return i
		}
	}
	return len(disk)
}

// Modifies `disk` in-place to move all file blocks as far left as possible.
func optimize(disk []int) {
	j := findFree(disk, 0)

	for i := len(disk) - 1; i >= 0; i-- {
		fid := disk[i]

		if fid == -1 {
			continue
		}

		j := findFree(disk, j)
		if j < i {
			// Can move
			disk[j] = fid
			disk[i] = -1
		}

		if j > i {
			// There's no way to optimize further
			break
		}
	}
}

// Calculates the checksum for `disk` - each blocks position multiplied by its file ID
// number
func checksum(disk []int) (sum int) {
	for i, fid := range disk {
		if fid == -1 {
			continue
		}

		sum += i * int(fid)
	}
	return
}

// Moves all file blocks one at a time as far left as possible in a representation
// of the filesystem blocks, then calculates a checksum.
func part1(bytes []int) int {
	totalSize := 0
	for _, b := range bytes {
		totalSize += int(b)
	}

	disk := make([]int, totalSize)
	di := 0
	fid := 0
	for i, b := range bytes {
		var x int
		if i % 2 == 0 {
			x = fid
		} else {
			x = -1
		}
		for dj := 0; dj < int(b); dj++ {
			disk[di+dj] = x
		}
		di += int(b)
		if x != -1 {
			fid++
		}
	}

	optimize(disk)
	return checksum(disk)
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	var bytes []int
	for line := range lines {
		// There's only one line.
		bytes = make([]int, len(line))
		for i, r := range []byte(line) {
			bytes[i] = int(r - 48)
		}
	}

	fmt.Println(part1(bytes))
}
