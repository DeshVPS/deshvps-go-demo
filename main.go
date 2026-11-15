// DeshVPS Go sample: a tiny JSON API on $PORT.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync/atomic"
)

var orders atomic.Int64

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<!doctype html><title>Cha Ghor Go</title><body style="font-family:system-ui;max-width:40rem;margin:4rem auto"><h1>Cha Ghor API, in Go</h1><p>Try <a href="/api/menu">/api/menu</a> or POST /api/orders.</p>`))
	})
	http.HandleFunc("/api/menu", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"go": runtime.Version(), "items": []map[string]any{{"name": "Malai cha", "taka": 40}, {"name": "Singara", "taka": 15}}})
	})
	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"order": orders.Add(1)})
	})
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
