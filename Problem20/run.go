package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	max := 100
	current := 1

	table := []int{1}
	overFlow := 0

	for current <= max {
		for i := len(table) - 1; i >= 0; i-- {
			value := table[i]
			value *= current
			if overFlow > 0 {
				value += overFlow
			}
			table[i] = value % 10
			if value >= 10 {
				if i == 0 {
					table = append([]int{(value % 100) / 10}, table...)
					if value >= 100 {
						table = append([]int{value / 100}, table...)
					}
				}
				overFlow = value / 10
			} else {
				overFlow = 0
			}

		}
		overFlow = 0
		current++
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