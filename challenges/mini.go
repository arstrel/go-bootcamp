package challenges

import (
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
