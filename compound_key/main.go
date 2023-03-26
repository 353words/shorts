/*
#golang gem: You can use struct as a compound key in maps.

type Point struct {
	X int
	Y int
}

colors := make(map[Point]int)
pt := Point{10, 20}
colors[pt] = 0xFF0000
fmt.Printf("%X\n", colors[pt]) // FF0000

*/

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	colors := make(map[Point]int)
	pt := Point{10, 20}
	colors[pt] = 0xFF0000
	fmt.Printf("%X\n", colors[pt])
}
