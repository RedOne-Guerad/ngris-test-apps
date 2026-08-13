package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

// Backend for the full-stack app: it only owns /api/* (backend_paths=["/api"]).
// Everything else is served as static files from the ngris edge.
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"message": "hello from the ngris backend", "path": r.URL.Path})
	})
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("ngris-fullstack backend listening on :%s (serves /api/*)", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
