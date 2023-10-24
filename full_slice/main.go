package main

import "fmt"

func main() {
	cart := []string{"orange", "apple", "bread"}
	fruit := cart[:2]
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit) // [orange apple lemon]
	fmt.Println("cart: ", cart)  // [orange apple lemon]

	cart = []string{"orange", "apple", "bread"}
	fruit = cart[:2:2]
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit) // [orange apple lemon]
	fmt.Println("cart: ", cart)  // [orange apple bread]
}
