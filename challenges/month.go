package challenges

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func DaysInMonth() {
	if len(os.Args) < 3 {
		fmt.Println("Give me a month name")
		return
	}

	// get the current year and find out whether it's a leap year
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// setting it to 28, saves me typing it below again
	days := 28

	month := os.Args[2]

	// case insensitive
	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		// try a "logical and operator" above.
		// like: `else if m == "february" && leap`
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}

func Seasons() {
	if len(os.Args) != 3 {
		fmt.Println("Gimme a month name.")
		return
	}

	// m := os.Args[1]

	// if m == "Dec" || m == "Jan" || m == "Feb" {
	// 	fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// 	fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jul" || m == "Aug" {
	// 	fmt.Println("Summer")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// 	fmt.Println("Fall")
	// } else {
	// 	fmt.Println("ERROR: not a month.")
	// }

	switch m := os.Args[2]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}
