package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	table := [21][21]int{}

	for r := 0; r < len(table); r++ {
		for c := 0; c < len(table[0]); c++ {
			if r == 0 || c == 0 {
				table[r][c] = 1
			} else {
				table[r][c] = table[r - 1][c] + table[r][c - 1]
			}
		}
	}

	result := table[len(table) - 1][len(table[0]) - 1]
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}