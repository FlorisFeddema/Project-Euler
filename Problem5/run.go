package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	result := 0

	current := 20

	for result == 0 {
		isCorrect := true
		for i := 2; i <= 20; i++  {
			if current % i != 0{
				isCorrect = false
				break
			}
		}
		if isCorrect {
			result = current
		}
		current += 20
	}

	fmt.Printf("Result: %[1]d\n", result)

}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}