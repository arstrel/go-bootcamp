package exerciseflowcontrol

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------
func SimpleIfs() {
	age, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}

}

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

func Simplify() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	var big bool

	if radius >= 200 {
		big = true
	}

	if big && isSphere {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func ArgCount() {
	args := os.Args[2:]

	switch len(args) {
	case 0:
		fmt.Println("Give me args")
	case 1:
		fmt.Printf("There is one: %s\n", args[0])
	case 2:
		fmt.Printf("There are two: %s\n", args[0]+" "+args[1])
	default:
		fmt.Print("There are ", len(args), " arguments\n")
	}
}

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func VowelOrConsonant() {
	args := os.Args[2:]

	if len(args) != 1 || len(args[0]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	letter := args[0]
	isVowel := strings.ContainsAny(letter, "aeiouAEIOU")

	if isVowel {
		fmt.Printf("%q is a vowel.\n", letter)
	} else if letter == "y" || letter == "w" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", letter)
	} else {
		fmt.Printf("%q is a consonant.\n", letter)
	}

}
