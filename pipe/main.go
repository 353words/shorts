/*
#golang #gem You can pipe output from one exec.Command to the other.

grep := exec.Command("grep", "func", "pipe.go")
out, err := grep.StdoutPipe()
if err != nil {
	return err
}
if err := grep.Start(); err != nil {
	return err
}
wc := exec.Command("wc", "-l")
wc.Stdin = out
data, err := wc.CombinedOutput()
*/
package main

import (
	"fmt"
	"os/exec"
)

func wcl() error {
	grep := exec.Command("grep", "func", "pipe.go")
	out, err := grep.StdoutPipe()
	if err != nil {
		return err
	}
	if err := grep.Start(); err != nil {
		return err
	}
	wc := exec.Command("wc", "-l")
	wc.Stdin = out
	data, err := wc.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	wcl()
}
