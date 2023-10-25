package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/agents", agentsHandler)

	addr := ":8080"
	log.Printf("info: server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func agentsHandler(w http.ResponseWriter, r *http.Request) {
	var wtr io.Writer = w
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gzw := gzip.NewWriter(w)
		defer gzw.Close()
		wtr = gzw
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(wtr)
	reply := map[string]any{
		"agents": db.QueryAgents(),
	}
	if err := enc.Encode(reply); err != nil {
		log.Printf("error: can't encode - %s", err)
	}
}

type Agent struct {
	ID     string
	Lat    float64
	Lng    float64
	Status string
}

var db DB

type DB struct{}

func (db DB) QueryAgents() []Agent {
	return []Agent{
		{"007", 51.4871871, -0.1270659, "free"},
		{"Q", 51.4871881, -0.1270759, "busy"},
		{"M", 51.4871878, -0.1270259, "free"},
	}
}
