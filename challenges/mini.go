package challenges

import (
	"fmt"
	"time"
)

func PartsOfDay() {
	switch now := time.Now().Hour(); {
	case now < 6:
		print("Good night!")
	case now < 12:
		print("Good morning!")
	case now < 18:
		print("Good afternoon!")
	default:
		print("Good evening!")

	}
}

func MultiplicationTable() {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if row < 1 {
				// print title of the table first
				if col < 1 {
					fmt.Printf("%5s", "X")
				} else {
					fmt.Printf("%5d", col)
				}
			} else {
				if col < 1 {
					fmt.Printf("%5d", row)
				} else {
					fmt.Printf("%5d", row*col)
				}
			}

		}
		print("\n")
	}
}
