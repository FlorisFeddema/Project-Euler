package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	target := 10001

	number, result, count:= 1, 0, 1
	for count <= target {
		number += 2
		isPrime := true

		if number % 2 == 0 {
			isPrime = false
		}
		for i := 3; i < number; i+=2{
			if number % i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++
			result = number
		}
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}