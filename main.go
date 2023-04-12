package main

import (
	"fmt"
	"os"

	"github.com/arstrel/go-bootcamp/strings"
	"github.com/arstrel/go-bootcamp/variables"
)

func main() {
	switch os.Args[1] {
	case "path":
		SeparatePath(os.Args[2])
	case "color":
		variables.Color()
	case "vars":
		variables.VarsToVars()
	case "assign":
		variables.Assign()
	case "assign2":
		variables.Assign2()
	case "multi":
		variables.MultiShort()
	case "swapper":
		variables.Swapper()
	case "swapper2":
		variables.Swapper2()
	case "discard":
		variables.DiscardFile()
	case "convert1":
		variables.Convert1()
	case "convert2":
		variables.Convert2()
	case "convert3":
		variables.Convert3()
	case "convert4":
		variables.Convert4()
	case "convert5":
		variables.Convert5()
	case "cli1":
		variables.CommandLine()
	case "cli2":
		variables.CommandLine2()
	case "cli3":
		variables.CommandLine3()
	case "cli4":
		variables.CommandLine4()
	case "banger":
		strings.Banger(os.Args[2])
	case "winpath":
		strings.WindowsPath()
	case "printjson":
		strings.PrintJSON()
	case "rawconcat":
		strings.RawConcat()
	case "countchars":
		strings.CountChars(os.Args[2])
	case "tolowercase":
		strings.ToLowercase()
	case "trimspace":
		strings.TrimSpace()
	case "trimright":
		strings.TrimRight()
	default:
		fmt.Println("Unknown operation")
	}
	fmt.Println("Done")

}
