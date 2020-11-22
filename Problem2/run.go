package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	valueCurrent := 2
	valueLast := 1
	sum := 0
	for valueCurrent <= 4000000 {
		if valueCurrent % 2 == 0{
			sum += valueCurrent
		}
		x := valueCurrent
		valueCurrent = valueCurrent + valueLast;
		valueLast = x
	}

	fmt.Printf("Result: %[1]d\n", sum)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s", elapsed)
}


