package exercise3

import (
	"fmt"
)

func Timezones() {
	const (
		EST = -(5 + iota)
		CST
		_ // blank identifier
		PST
	)
	fmt.Println(EST, CST, PST)
}

// ---------------------------------------------------------
// EXERCISE: Iota Months
//
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
//
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
//
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------

func Months() {
	const (
		Sep = iota + 9
		Oct
		Nov
	)

	fmt.Println(Sep, Oct, Nov)
}

// ---------------------------------------------------------
// EXERCISE: Iota Months #2
//
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------

func Months2() {
	// HINT: This is a valid constant declaration
	//       Blank-Identifier can be used in place of a name
	const _ = iota
	//    ^- this is just a name

	// Now, use iota and initialize the following constants
	// "automatically" to 1, 2, and 3 respectively.
	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

func Seasons() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
