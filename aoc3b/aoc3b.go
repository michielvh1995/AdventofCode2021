package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Let's fucking do this: Radix sort, baby!
func RadixSortOnes(inputArray [][]int, digitOfSignificance int) []int {
	// Radix sort works by taking the first bit in the list, and sorting all arrays on them.
	// We will use a list; we push the 1s to the front, and the 0s to the back!

	if len(inputArray) > 1 {
		// Alright.. we'll sort them into two lists, as we only need to most/least common one.
		onesArray := make([][]int, 0)
		zeroesArray := make([][]int, 0)

		for i := 0; i < len(inputArray); i++ {
			// We radixsort the list by pushing the items to the front & back of the list.
			if inputArray[i][digitOfSignificance] == 1 {
				onesArray = append(onesArray, inputArray[i])
			} else {
				zeroesArray = append(zeroesArray, inputArray[i])
			}
		}

		if len(onesArray) >= len(zeroesArray) {
			return RadixSortOnes(onesArray, digitOfSignificance+1)
		} else {
			return RadixSortOnes(zeroesArray, digitOfSignificance+1)
		}
	}
	return inputArray[0]
}

// Let's fucking do this: Radix sort, baby!
func RadixSortZeroesPreference(inputArray [][]int, digitOfSignificance int) []int {
	// Radix sort works by taking the first bit in the list, and sorting all arrays on them.
	// We will use a list; we push the 1s to the front, and the 0s to the back!

	if len(inputArray) > 1 {
		// Alright.. we'll sort them into two lists, as we only need to most/least common one.
		onesArray := make([][]int, 0)
		zeroesArray := make([][]int, 0)

		for i := 0; i < len(inputArray); i++ {
			// We radixsort the list by pushing the items to the front & back of the list.
			if inputArray[i][digitOfSignificance] == 1 {
				onesArray = append(onesArray, inputArray[i])
			} else {
				zeroesArray = append(zeroesArray, inputArray[i])
			}
		}

		if len(onesArray) >= len(zeroesArray) {
			return RadixSortZeroesPreference(zeroesArray, digitOfSignificance+1)
		} else {
			return RadixSortZeroesPreference(onesArray, digitOfSignificance+1)

		}
	}
	return inputArray[0]
}

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

func ParseFile(r io.Reader) ([][]int, error) {
	// We use this function to parse the input file.
	scanner := bufio.NewScanner(r)

	// We get here the width of the array
	var sz int

	if scanner.Scan() {
		sz = len(scanner.Text())
	}

	// Here's the output array:
	fileContents := make([][]int, 0)
	fileContents = append(fileContents, ParseLine(scanner.Text(), sz))

	// Then we read the rest of the file line by line
	for scanner.Scan() {
		var line = ParseLine(scanner.Text(), sz)
		fileContents = append(fileContents, line)
	}

	return fileContents, scanner.Err()
}

// This function turns an array of 1s and 0s into an integer
func BinaryIntArrayToInt(binaryArray []int) int {
	var outputInt = 0
	fmt.Println(binaryArray)

	for i := 0; i < len(binaryArray); i++ {
		// We bitshift this binary value with its position in the array
		outputInt += binaryArray[i] << (len(binaryArray) - i - 1)
	}

	return outputInt
}

func main() {
	// Get the input file handle:
	f, err := os.Open("./input3.txt")
	check(err)

	// Defer delays this function until the current function returns.
	// So in this case; the file handle will be closed when main() returns.
	defer f.Close()

	fileContents, err := ParseFile(f)
	check(err)

	// We calculate the O2 generator * CO2 scrubber
	OTwoValue := RadixSortOnes(fileContents, 0)
	COTwoValue := RadixSortZeroesPreference(fileContents, 0)

	// FInally turn it into an int and print it
	fmt.Println(BinaryIntArrayToInt(OTwoValue) * BinaryIntArrayToInt(COTwoValue))
}

func check(err error) {
	// Quit the program if there is an error
	if err != nil {
		panic(err)
	}
}
