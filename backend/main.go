package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok", "service": "ngris-backend-test"})
	})
	mux.HandleFunc("/api/echo", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"method": r.Method, "path": r.URL.Path, "host": r.Host})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ngris backend test v2 running in-cluster. Try /api/health or /api/echo\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("ngris-backend-test listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
