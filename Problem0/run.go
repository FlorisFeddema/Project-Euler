package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	//TODO: MAKE CODE

	result := 0
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}