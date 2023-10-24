package main

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dataHandler)

	addr := ":8080"
	log.Printf("info: server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	var wtr io.Writer = w
	if r.Header.Get("Accept-Encoding") == "gzip" {
		w.Header().Set("Content-Encoding", "gzip")
		gzw := gzip.NewWriter(w)
		defer gzw.Close()
		wtr = gzw
	}

	w.Header().Set("Content-Type", "text/csv")
	cw := csv.NewWriter(wtr)
	defer cw.Flush()
	for _, row := range query() {
		cw.Write(row)
	}
}

func query() [][]string {
	return [][]string{
		{"agent", "lat", "lng", "status"},
		{"007", "1", "2", "free"},
		{"Q", "3", "4", "busy"},
		{"M", "5", "6", "free"},
	}
}
