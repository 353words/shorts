/*
#golang #gem: Use your own HTTP router and server to avoid third party packages injecting handlers.

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/debug/vars", expvar.Handler())

	srv := http.Server{
	    Addr:    ":8080",
	    Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
	    log.Fatalf("error: %s", err)
	}
*/
package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/debug/vars", expvar.Handler())

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
