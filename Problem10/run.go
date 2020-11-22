package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	current := 3
	max := 2000000
	sum := 2

	for current < max {
		isPrime := true

		sqrt := 0
		for i := 1; i <= current; i++ {
			sq := i*i
			if sq > current {
				sqrt = i
				break
			}
		}

		for i := 3; i < sqrt; i+=2{
			if current % i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			sum += current
		}
		current += 2
	}
	result := sum
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}