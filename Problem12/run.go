package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	maxDivisors := 500
	currentDivisors := 0

	currentTiangle := 1
	currentValue := 0

	for currentDivisors <= maxDivisors {
		currentValue += currentTiangle;

		sqrt := 0
		for i := 1; i <= currentValue; i++ {
			sq := i*i
			if sq > currentValue {
				sqrt = i
				break
			}
		}

		currentDivisors = 1
		for i := 1; i <= sqrt ; i++ {
			if currentValue % i == 0 {
				if currentValue/ i == i {
					currentDivisors++
				} else {
					currentDivisors += 2
				}
			}
		}

		currentTiangle++
	}


	result := currentValue
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}