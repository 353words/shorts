/*
#golang gem: Use strings.EqualFold to do Unicode aware case insensitive comparison

s1, s2 := "Σ", "ς"

fmt.Println(strings.ToLower(s1) == strings.ToLower(s2)) // false
fmt.Println(strings.EqualFold(s1, s2))                  // true
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1, s2 := "Σ", "ς"

	fmt.Println(strings.ToLower(s1) == strings.ToLower(s2)) // false
	fmt.Println(strings.EqualFold(s1, s2))                  // true
}
