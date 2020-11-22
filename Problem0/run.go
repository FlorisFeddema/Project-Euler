package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s", elapsed)
}