package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "I ♡ Go"
	fmt.Println(len(text)) // 8
	n := utf8.RuneCountInString(text)
	fmt.Println(n) // 6
}
