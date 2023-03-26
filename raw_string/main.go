/*
#golang gem: Use raw strings to escape \ or create multiline strings

s := `C:\to\be\or\not\to.bee`
re := regexp.MustCompile(`\d+`)
poem := `
The Road goes ever on and on,
Down from the door where it began.
`
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := `C:\to\be\or\not\to.bee`
	fmt.Println(s)
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re)
	poem := `
	The Road goes ever on and on,
	Down from the door where it began.
	`
	fmt.Println(poem)
}
