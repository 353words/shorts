/*
#golang #gem: fmt.Printf can access arguments by location.

	name, age := "Bugs", 82
	fmt.Printf(
		"%[1]s is %[2]d years old. %[2]d!\n",
		name, age,
	)
	// Bugs is 82 years old. 82!
*/
package main

import (
	"fmt"
)

func main() {
	name, age := "Bugs", 82
	fmt.Printf(
		"%[1]s is %[2]d years old. %[2]d!\n",
		name, age,
	)
	// Bugs is 82 years old. 82!
}
