package input

import (
	"bufio"
	"fmt"
	"os"
)


// A generator that produces a stream of text lines by reading the given file.
// Returns nil, error if there was an error opening the file.
// Otherwise writes lines of text to `lines`.
func ReadLines(name string) (lines chan string, err error) {
	file, err := os.Open(name)

	if err != nil {
		return
	}

	lines = make(chan string)

	go func() {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	return
}
