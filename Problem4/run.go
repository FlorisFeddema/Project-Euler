package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	max := 999

	x := max
	y := max

	result := 0

	var out = ""

	for x > 0 || y > 0 {
		if x == 0  {
			y--
			x = y
		} else {
			x--
		}
		out = strconv.Itoa(x * y)

		isPalindrome := true
		for i, _ := range out {
			if out[i] != out[len(out) - i - 1]{
				isPalindrome = false
			}
		}
		if isPalindrome && x * y > result {
			result = x * y
		}
	}

	fmt.Printf("Result : %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}