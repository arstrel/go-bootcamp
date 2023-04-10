package variables

import (
	"fmt"
	"log"
	"path"
)

// ---------------------------------------------------------
// EXERCISE: Make It Blue
//
//  1. Change `color` variable's value to "blue"
//
//  2. Print it
//
// EXPECTED OUTPUT
//  blue
// ---------------------------------------------------------

func Color() {
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW:

	color = "blue"

	log.Println(color)
}

// ---------------------------------------------------------
// EXERCISE: Variables To Variables
//
//  1. Change the value of `color` variable to "dark green"
//
//  2. Do not assign "dark green" to `color` directly
//
//     Instead, use the `color` variable again
//     while assigning to `color`
//
//  3. Print it
//
// RESTRICTIONS
//  WRONG ANSWER, DO NOT DO THIS:
//  `color = "dark green"`
//
// HINT
//  + operator can concatenate string values
//
//  Example:
//
//  "h" + "e" + "y" returns "hey"
//
// EXPECTED OUTPUT
//  dark green
// ---------------------------------------------------------

func VarsToVars() {
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW

	color = "dark " + color

	// UNCOMMENT THE CODE BELOW TO PRINT THE VARIABLE

	log.Println(color)
}

// ---------------------------------------------------------
// EXERCISE: Assign With Expressions
//
//  1. Multiply 3.14 with 2 and assign it to `n` variable
//
//  2. Print the `n` variable
//
// HINT
//  Example: 3 * 2 = 6
//  * is the product operator (it multiplies numbers)
//
// EXPECTED OUTPUT
//  6.28
// ---------------------------------------------------------

func Assign() {
	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0
	n := 0.

	// ADD YOUR CODE BELOW

	n = 3.14 * 2

	log.Println(n)
}

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func Assign2() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet = "Mars"
	isTrue = true
	temp = 19.5

	fmt.Printf("Air is good on %v\n", planet)
	fmt.Printf("It's %v\n", isTrue)
	fmt.Printf("It is %v degrees\n", temp)
}

// ---------------------------------------------------------
// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable
//
// NOTE
//  You should use `multi` function
//  while initializing the variables
//
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------

func MultiShort() {
	// ADD YOUR DECLARATIONS HERE

	_, b := multi()

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}

// EXERCISE: Swapper
//
//  1. Change `color` to "orange"
//     and `color2` to "green" at the same time
//
//     (use multiple-assignment)
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------

func Swapper() {
	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

func Swapper2() {
	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"
	red, blue = blue, red

	fmt.Println(red, blue)
}

// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

func DiscardFile() {
	// UNCOMMENT THE CODE BELOW:

	dir, _ := path.Split("secret/file.txt")

	fmt.Println(dir)

}
