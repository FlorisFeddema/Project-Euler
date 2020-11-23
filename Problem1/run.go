package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	sum := 0

	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}		
	}
	
	fmt.Printf("Result: %[1]d\n", sum)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}
