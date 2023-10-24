package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Run actual checks.
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/api/health", healthHandler)

	addr := ":8080"
	log.Printf("info: server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Printf("error: %s", err)
	}
}
