package exercise2

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------
func CountChars(s string) int {
	l := utf8.RuneCountInString(s)
	fmt.Printf("Length of string is %v. It takes %v bytes", l, len(s))
	return l
}

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------
func Banger(s string) string {
	l := utf8.RuneCountInString(s)
	bangs := strings.Repeat("!", l)
	c := strings.ToUpper(s)
	fmt.Println(bangs + c + bangs)
	return bangs + c + bangs
}

// ---------------------------------------------------------
// EXERCISE: Windows Path
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func WindowsPath() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"

	p := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)
	fmt.Println()
	fmt.Println(p)
}

// ---------------------------------------------------------
// EXERCISE: Print JSON
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func PrintJSON() {
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"
	j := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
	fmt.Println(j)
}

// ---------------------------------------------------------
// EXERCISE: Raw Concat
//
//  1. Initialize the name variable
//     by getting input from the command line
//
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
//
// NOTE
//  You should concatenate the name variable in the correct
//  place.
//
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func RawConcat() {
	// uncomment the code below
	name := os.Args[2]

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := `hi ` + name + `!` + `
how are you?`

	fmt.Println(msg)
}

// ---------------------------------------------------------
// EXERCISE: ToLowercase
//
//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
//
// HINT
//  Check out the strings package from Go online documentation.
//  You will find the lowercase function there.
//
// INPUT
//  "SHEPARD"
//
// EXPECTED OUTPUT
//  shepard
// ---------------------------------------------------------
func ToLowercase() {
	name := os.Args[2]
	fmt.Println(strings.ToLower(name))
}

// ---------------------------------------------------------
// EXERCISE: Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     the given string
//  3. Trim the text variable and print it
//
// EXPECTED OUTPUT
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------

func TrimSpace() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func TrimRight() {
	// currently it prints 17
	// it should print 5

	name := "inanç           "
	trimmed := strings.TrimRight(name, " ")
	l := utf8.RuneCountInString(trimmed)
	fmt.Println(l)
}
