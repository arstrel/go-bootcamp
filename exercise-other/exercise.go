package exerciseother

import (
	"fmt"
	"os"
	"strconv"
)

func Other1() {
	// UNCOMMENT THIS TO SEE IT IN ACTION:
	var n int

	if a := os.Args[1:]; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// else if here shadows the main func's n
		// variable and it declares it own
		// with the same name: n

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// n here belongs to the main func
	// not to the if statement above

	// UNCOMMENT ALSO LINES BELOW TO SEE IT IN ACTION:
	fmt.Printf("n is %d. 👻 👻 👻 - you've been shadowed ;-)\n", n)
}

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------
func OddOrEven() {
	if len(os.Args) < 3 {
		fmt.Println("Pick a number")
		return
	}
	args := os.Args[1:]
	if n, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println(err)
	} else if n%2 == 0 && n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}
}
