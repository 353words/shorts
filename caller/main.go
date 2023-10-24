package main

import (
	"fmt"
	"log"
	"runtime"
)

func debug(format string, args ...interface{}) {
	_, file, lnum, ok := runtime.Caller(1)
	if !ok {
		file, lnum = "???", -1
	}
	msg := fmt.Sprintf(format, args...)
	log.Printf("%s:%d: %s\n", file, lnum, msg)
}

func inc(n int) int {
	debug("inc(%v)", n)
	return n + 1
}

func main() {
	inc(352)
}