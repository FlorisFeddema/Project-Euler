package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	result := 0

	max := 10000
	current := 0

	var total []int

	for current <= max {
		sum := 0
		for i := 1; i < current; i++ {
			if current % i == 0 {
				sum += i
			}
		}
		total = append(total, sum)
		current++
	}

	for index, element := range total {
		if element <= len(total) {
			if index == total[element] && element != index {
				result += index
			}
		}
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}