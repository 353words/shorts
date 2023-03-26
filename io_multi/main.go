/*
#golang gem: You can use io.MultiWriter to write to several destinations

w := io.MultiWriter(file, os.Stdout)
logger := log.New(w, "app: ", log.LstdFlags)
logger.Printf("INFO: %q logged in", "joe")
*/
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("/tmp/app.log", flag, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := io.MultiWriter(file, os.Stdout)
	logger := log.New(w, "app: ", log.LstdFlags)
	logger.Printf("INFO: %q logged in", "joe")
	// app: 2020/10/07 14:00:11 INFO: "joe" logged in
}
