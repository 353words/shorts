/*
#golang #gem: Send signal 0 to check if a process is alive.

func isAlive(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}

func main() {
	fmt.Println(isAlive(os.Getpid())) // true
	fmt.Println(isAlive(666))         // false
}
*/
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println(isAlive(os.Getpid())) // true
	fmt.Println(isAlive(666))         // false
}

func isAlive(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}

/*
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := waitTermination(ctx, 45273); err != nil {
		log.Fatal(err)
	}
}

func waitTermination(ctx context.Context, pid int) error {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			p, err := os.FindProcess(pid)
			if err != nil {
				return nil
			}
			fmt.Println(p)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
*/
