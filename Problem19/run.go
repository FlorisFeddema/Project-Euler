package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	endYear := 2000
	result := 0

	dayOfWeek := 1
	year := 1901
	month := 0
	day := 0

	for year <= endYear {
		for month < 12 {
			for day < getDaysMonth(month, year) {
				if dayOfWeek == 6 && day == 0 {
					result ++
				}
				dayOfWeek++
				dayOfWeek = dayOfWeek % 7
				day++
			}
			day = 0
			month++
		}
		month = 0
		year++
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func getDaysMonth(currentMonth int, currentYear int) int {
	if currentMonth == 1 {
		if currentYear % 400 == 0 {
			return 29
		}
		if currentYear % 100 == 0 {
			return 28
		}
		if currentYear % 4 == 0 {
			return 29
		}
	}
	months := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	return months[currentMonth]
}


func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}