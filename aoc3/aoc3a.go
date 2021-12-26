package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// This function parses the line of input,and returns it as an array.
func ParseLine(line string, size int) []int {
	// we loop over all the characters in the string
	//  looping over strings, c-style is a bit weird in Go; it uses UTF-8 default format
	// So one and two byte long chars.
	// But since we only have 1s and 0s it should be fine.

	var outArray = make([]int, size)

	// we parse the string into an array of 1 and 0 ints.
	for i := 0; i < len(line); i++ {
		if line[i] == '1' {
			outArray[i] = 1
		} else {
			outArray[i] = 0
		}
	}

	return outArray
}

func CalculateDepth(r io.Reader) (int, error) {
	// Read the lines one by one:
	//  For this we first need to create a bufio scanner
	scanner := bufio.NewScanner(r)

	// We get here the width of the array
	var sz, count int
	if scanner.Scan() {
		sz = len(scanner.Text())
	}

	var bitCounts = ParseLine(scanner.Text(), sz)

	// Then we read the rest of the file line by line
	for scanner.Scan() {
		var line = ParseLine(scanner.Text(), sz)
		count += 1

		// Add the arrays (aka increase the bits by atmost 1)
		for i := 0; i < sz; i++ {
			bitCounts[i] += line[i]
		}
	}

	// Now we calculate the epsilon and gamma values
	var epsilon = make([]int, sz)
	var gamma = make([]int, sz)

	// Gamma is the most common, epsilon the least common
	for i := 0; i < sz; i++ {
		if bitCounts[i] < count/2 {
			epsilon[i] = 1
			gamma[i] = 0
		} else {
			epsilon[i] = 0
			gamma[i] = 1
		}
	}

	// Now turn the gamma and epsilon into ints, and multiply them
	var eps = strings.Trim(strings.Replace(fmt.Sprint(epsilon), " ", "", -1), "[]")
	var gam = strings.Trim(strings.Replace(fmt.Sprint(gamma), " ", "", -1), "[]")

	var epsVal int
	var gamVal int

	if epv, err := strconv.ParseInt(eps, 2, 64); err != nil {
		panic(err)
	} else {
		epsVal = int(epv)
	}

	if gamv, err := strconv.ParseInt(gam, 2, 64); err != nil {
		panic(err)
	} else {
		gamVal = int(gamv)
	}

	return gamVal * epsVal, scanner.Err()
}

func main() {
	// Get the input file handle:
	f, err := os.Open("./input3.txt")
	check(err)

	// Defer delays this function until the current function returns.
	// So in this case; the file handle will be closed when main() returns.
	defer f.Close()

	fmt.Println(CalculateDepth(f))
}

func check(err error) {
	// Quit the program if there is an error
	if err != nil {
		panic(err)
	}
}
