package exerciseprint

import "fmt"

func PrintBio() {
	var (
		firstName = "Artem"
		lastName  = "Strel"
		age       = 35
	)
	tf := true
	temp := 29.5
	cl := "Hello world!"

	fmt.Printf("Hi! My Name is %s %s.\nI'm %d years old\n", firstName, lastName, age)
	fmt.Printf("That's %t.\nAnd the type is %[1]T \n", tf)
	fmt.Printf("The temperature is %.1f\n", temp)
	fmt.Printf("The type of temperature var is %T\n", temp)
	fmt.Printf("And the classics: \"%s\"\n", cl)

}
