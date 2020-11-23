package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	max := 600851475143

	current := 1

	for current < max {
		if max % current == 0 {
			max = max / current
		}
		current++
	}
	fmt.Printf("Result: %[1]d\n", max)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}