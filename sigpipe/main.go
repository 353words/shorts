package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGPIPE)
	go func() {
		<-ch
		os.Exit(0)
	}()

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

// $ go run seq.go | head
