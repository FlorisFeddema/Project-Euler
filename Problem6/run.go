package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	sum, square := 0, 0

	for i := 1; i < 101; i++ {
		sum += i
		square += i*i
	}
	sum *= sum

	result := sum - square
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}