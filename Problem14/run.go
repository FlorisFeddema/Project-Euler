package main

import (
	"fmt"
	"time"
)

var maxValue int
var longestChain int
var longestNumber int
var currentNumber int
var currentChain int

func main() {
	defer timeTrack(time.Now())

	maxValue = 1000000
	longestChain = 0
	longestNumber = 0
	currentNumber = 2

	for currentNumber < maxValue {
		currentChain = 1

		solve(currentNumber)

		if currentChain > longestChain {
			longestChain = currentChain
			longestNumber = currentNumber
		}

		currentNumber++
	}

	result := longestNumber
	fmt.Printf("Result: %[1]d\n", result)
}

func solve(value int) int {
	currentChain++
	if value == 1 {
		return 1
	}
	if value % 2 == 0 {
		value = value / 2
	} else {
		value = 3*value + 1
	}
	return solve(value);
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}