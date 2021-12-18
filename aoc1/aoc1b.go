package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func CalculateDepth(r io.Reader) (int, error) {
	// Read the lines one by one:
	//  For this we first need to create a bufio scanner
	scanner := bufio.NewScanner(r)
	var counts int // amount of increases

	a, b, c := 10000, 10000, 10000

	// The we read the file line by line
	for scanner.Scan() {
		// Here we do the str to int conversion
		depth, err := strconv.Atoi(scanner.Text())
		check(err)

		// If the frame of the latest is deeper than the previous one
		if depth+b+c > a+b+c {
			counts++
		}

		// Now pass on the values
		a = b
		b = c
		c = depth
	}

	return counts, scanner.Err()
}

func main() {
	// Get the input file handle:
	f, err := os.Open("./input.txt")
	check(err)

	// Defer delays this function until the current function returns.
	// So in this case; the file handle will be closed when main() returns.
	defer f.Close()

	fmt.Println(CalculateDepth(f))
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
