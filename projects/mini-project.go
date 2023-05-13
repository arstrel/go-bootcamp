package projects

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Randomization() {
	rand.Seed(time.Now().UnixNano())

	guess := 10
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
}

func Guess() {
	param := os.Args[2:]

	if len(param) != 1 {
		fmt.Println("Please provide a number.")
		return
	}

	uinput, err := strconv.Atoi(param[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	guess := rand.Intn(uinput)

	if uinput == guess {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}

}
