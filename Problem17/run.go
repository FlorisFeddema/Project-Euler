package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	current := 1
	max := 1000

	sum := 0
	andLarge := false
	andSmall := false
	ten := false

	for current <= max {
		andLarge = false
		andSmall = false
		ten = false

		rest := current
		if rest >= 1000 {
			digit := current/1000
			if digit == 1 {
				sum += 3 // one
			}
			sum += 8 // thousand
			andLarge = true
		}
		rest = rest % 1000
		if rest >= 100 {
			digit := rest/100
			if digit == 1 {
				sum += 3 // one
			} else if digit == 2 {
				sum += 3 // two
			} else if digit == 3 {
				sum += 5 // three
			} else if digit == 4 {
				sum += 4 // four
			} else if digit == 5 {
				sum += 4 // five
			} else if digit == 6 {
				sum += 3 // six
			} else if digit == 7 {
				sum += 5 // seven
			} else if digit == 8 {
				sum += 5 // eight
			} else if digit == 9 {
				sum += 4 // nine
			}
			sum += 7 // hundred
			andLarge = true
		}
		rest = rest % 100
		if rest >= 10 {
			digit := rest/10
			if digit >= 2 {
				if digit == 2 {
					sum += 6 // twenty
				} else if digit == 3 {
					sum += 6 // thirty
				} else if digit == 4 {
					sum += 5 // forty
				} else if digit == 5 {
					sum += 5 // fifty
				} else if digit == 6 {
					sum += 5 // sixty
				} else if digit == 7 {
					sum += 7 // seventy
				} else if digit == 8 {
					sum += 6 // eighty
				} else if digit == 9 {
					sum += 6 // ninety
				}
			} else if digit == 1 {
				digit = rest % 10
				if digit == 0 {
					sum += 3 // ten
				} else if digit == 1 {
					sum += 6 // eleven
				} else if digit == 2 {
					sum += 6 // twelve
				} else if digit == 3 {
					sum += 8 // thirteen
				} else if digit == 4 {
					sum += 8 // fourteen
				} else if digit == 5 {
					sum += 7 // fifteen
				} else if digit == 6 {
					sum += 7 // sixteen
				} else if digit == 7 {
					sum += 9 // seventeen
				} else if digit == 8 {
					sum += 8 // eighteen
				} else if digit == 9 {
					sum += 8 // nineteen
				}
				ten = true
			}
			andSmall = true
		}
		if !ten {
			rest = rest % 10
			digit  := rest
			if digit == 1 {
				sum += 3 // one
			} else if digit == 2 {
				sum += 3 // two
			} else if digit == 3 {
				sum += 5 // three
			} else if digit == 4 {
				sum += 4 // four
			} else if digit == 5 {
				sum += 4 // five
			} else if digit == 6 {
				sum += 3 // six
			} else if digit == 7 {
				sum += 5 // seven
			} else if digit == 8 {
				sum += 5 // eight
			} else if digit == 9 {
				sum += 4 // nine
			}
			if digit != 0 {
				andSmall = true
			}
		}

		if andLarge && andSmall {
			sum += 3 //and
		}

		current++
	}

	result := sum
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}