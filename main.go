package main

import (
	"fmt"
	"os"

	challenges "github.com/arstrel/go-bootcamp/challenges"
	exerciseflowcontrol "github.com/arstrel/go-bootcamp/exercise-flowcontrol"
	exercise3 "github.com/arstrel/go-bootcamp/exercise-iota"
	exerciseother "github.com/arstrel/go-bootcamp/exercise-other"
	exerciseprint "github.com/arstrel/go-bootcamp/exercise-print"
	exercise2 "github.com/arstrel/go-bootcamp/exercise-strings"
	exercise1 "github.com/arstrel/go-bootcamp/exercise-variables"
)

func main() {
	switch os.Args[1] {
	case "path":
		SeparatePath(os.Args[2])
	case "color":
		exercise1.Color()
	case "vars":
		exercise1.VarsToVars()
	case "assign":
		exercise1.Assign()
	case "assign2":
		exercise1.Assign2()
	case "multi":
		exercise1.MultiShort()
	case "swapper":
		exercise1.Swapper()
	case "swapper2":
		exercise1.Swapper2()
	case "discard":
		exercise1.DiscardFile()
	case "convert1":
		exercise1.Convert1()
	case "convert2":
		exercise1.Convert2()
	case "convert3":
		exercise1.Convert3()
	case "convert4":
		exercise1.Convert4()
	case "convert5":
		exercise1.Convert5()
	case "cli1":
		exercise1.CommandLine()
	case "cli2":
		exercise1.CommandLine2()
	case "cli3":
		exercise1.CommandLine3()
	case "cli4":
		exercise1.CommandLine4()
	case "banger":
		exercise2.Banger(os.Args[2])
	case "winpath":
		exercise2.WindowsPath()
	case "printjson":
		exercise2.PrintJSON()
	case "rawconcat":
		exercise2.RawConcat()
	case "countchars":
		exercise2.CountChars(os.Args[2])
	case "tolowercase":
		exercise2.ToLowercase()
	case "trimspace":
		exercise2.TrimSpace()
	case "trimright":
		exercise2.TrimRight()
	case "timezones":
		exercise3.Timezones()
	case "months":
		exercise3.Months()
	case "months2":
		exercise3.Months2()
	case "seasons-iota":
		exercise3.Seasons()
	case "bio":
		exerciseprint.PrintBio()
	case "ifs1":
		exerciseflowcontrol.SimpleIfs()
	case "ifs2":
		exerciseflowcontrol.Simplify()
	case "ifs3":
		exerciseflowcontrol.ArgCount()
	case "ifs4":
		exerciseflowcontrol.VowelOrConsonant()
	case "challenge1":
		challenges.Challenge1()
	case "challenge2":
		challenges.Challenge2()
	case "challenge3":
		challenges.Challenge3()
	case "movie":
		challenges.MoviePGRatingChecker()
	case "other1":
		exerciseother.Other1()
	case "oddoreven":
		exerciseother.OddOrEven()
	case "daysinmonth":
		challenges.DaysInMonth()
	case "seasons":
		challenges.Seasons()
	default:
		fmt.Println("Unknown operation")
	}
	fmt.Println("Done")

}
