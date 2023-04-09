package main

import (
	"fmt"
	"os"
)

func main() {
	op := os.Args[1]
	switch op {
	case "path":
		SeparatePath(os.Args[2])
	}
	fmt.Println("Done")

}
