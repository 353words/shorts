package main

import (
	"embed"
	"log"
	"net/http"
)

var (
	//go:embed static
	static embed.FS
)

func main() {
	handler := http.FileServer(http.FS(static))
	http.Handle("/static/", handler)

	addr := ":8080"
	log.Printf("info: server starting on %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
