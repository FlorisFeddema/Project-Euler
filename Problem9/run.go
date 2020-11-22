package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	a, b, c:= 1, 1, 1
	asq, bsq, csq := 0, 0, 0
	answer := 1000

	for b <= answer {
		if a == answer	{
			b++
			a = 0
		} else {
			a++
		}

		asq = a*a
		bsq = b*b
		csq = 0
		leftTotal := asq + bsq

		for i := 1; i <= leftTotal; i++ {
			sq := i*i
			if sq == leftTotal {
				c = i
				csq = sq
				break
			}
			if sq > leftTotal {
				break
			}
		}

		if a + b + c == answer && asq + bsq == csq {
			break
		}

	}

	result := a * b * c
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}