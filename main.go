package main

import (
	"fmt"
	"os"

	projects "github.com/arstrel/go-bootcamp/projects"
)

func main() {
	switch os.Args[1] {
	case "guess":
		projects.Guess()
	default:
		fmt.Println("Unknown operation")
	}
	fmt.Println("\nDone")

}
