package main

import (
	"fmt"
	"runtime"
)

func relu(n int) int {
	if n < 0 {
		runtime.Breakpoint()
		return 0
	}
	return n
}

func main() {
	n := relu(-7)
	fmt.Println(n)
}
