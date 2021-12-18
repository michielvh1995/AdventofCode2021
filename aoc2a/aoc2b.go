package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CalculateDepth(r io.Reader) (int, error) {
	// Read the lines one by one:
	//  For this we first need to create a bufio scanner
	scanner := bufio.NewScanner(r)

	var x, y, aim int

	// The we read the file line by line
	for scanner.Scan() {
		// Here we do the str to int conversion
		command := strings.Split(scanner.Text(), " ")
		movement, err := strconv.Atoi(command[1])
		check(err)

		// Now we turn the second part into the amount
		if command[0] == "forward" {
			x += movement
			y += movement * aim
		} else if command[0] == "down" {
			aim += movement
		} else if command[0] == "up" {
			aim -= movement
		}

	}

	return x * y, scanner.Err()
}

func main() {
	// Get the input file handle:
	f, err := os.Open("./input2a.txt")
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
