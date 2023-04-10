package variables

import (
	"fmt"
	"os"
	"path"
)

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func CommandLine() {
	// UNCOMMENT & FIX THIS CODE
	count := len(os.Args[1:])

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)
}

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------

func CommandLine2() {
	dir, _ := path.Split(os.Args[0])

	fmt.Println(dir)

}

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func CommandLine3() {
	if len(os.Args) <= 2 {
		fmt.Println("Please provide your name")
		return
	}
	name := os.Args[2]

	fmt.Printf("Hi %s\nHow are you?\n", name)

}

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------
func CommandLine4() {
	if len(os.Args[1:]) <= 5 {
		fmt.Println("Please provide 5 names")
		return
	}
	fmt.Println("There are 5 people!")
	for _, name := range os.Args[2:] {
		fmt.Printf("Hello great %s!\n", name)
	}
	fmt.Println("Nice to meet you all!")
}
