package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	table := []int{2}
	iterations := 1000
	currentIteration := 1
	overFlow := false

	for currentIteration < iterations {
		for i := len(table) - 1; i >= 0; i-- {
			value := table[i]
			value *= 2
			if overFlow {
				value++
			}
			table[i] = value % 10
			if value >= 10 {
				value -= 10
				if i == 0 {
					table = append([]int{1}, table...)
				} else {
					table[i - 1] = table[i -1]
				}
				overFlow = true
			} else {
				overFlow = false
			}

		}
		overFlow = false
		currentIteration++
	}


	result := 0
	for i := 0; i < len(table); i++ {
		result += table[i]
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}