package challenges

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func Challenge1() {
	var (
		validName = "jack"
		validPass = "1888"
	)
	args := os.Args[2:]

	if len(args) != 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	username, password := args[0], args[1]

	if username != validName {
		fmt.Printf("Access denied for \"%s\"\n", username)

	} else if password != validPass {
		fmt.Printf("Invalid password for \"%s\"\n", username)

	} else {
		fmt.Printf("Access granted to \"%s\"\n", username)
	}
}

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

type user struct {
	name string
	pass string
}

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
)

var users = []user{
	{"jack", "1888"},
	{"inanc", "1879"},
}

func checkUser(u, p string) string {
	for _, user := range users {
		if u == user.name {
			if p == user.pass {
				return fmt.Sprintf(accessOK, u)
			} else {
				return fmt.Sprintf(errPwd, u)
			}
		}
	}
	return fmt.Sprintf(errUser, u)
}

func Challenge2() {
	args := os.Args[2:]

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	u, p := args[0], args[1]

	res := checkUser(u, p)
	fmt.Println(res)
}

// ---------------------------------------------------------
// CHALLENGE
//  Add error handling to the feet to meters program.
//
// EXPECTED OUTPUT
//  go run main.go hello
//    error: 'hello' is not a number.
//
//  go run main.go what
//    error: 'what' is not a number.
//
//  go run main.go 100
//    100 feet is 30.48 meters.
// ---------------------------------------------------------

const template = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

func Challenge3() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(template))
		return
	}

	arg := os.Args[2]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
