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
		w.Write([]byte(`<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Cha Ghor API</title>
<style>*{box-sizing:border-box}body{margin:0;font-family:system-ui,-apple-system,"Segoe UI",Roboto,sans-serif;color:#1f2937;background:#fffbf5}
header{max-width:960px;margin:0 auto;padding:20px 24px;font-size:22px;font-weight:800;color:#065f46}
.hero{max-width:960px;margin:12px auto;padding:48px 32px;border-radius:24px;background:linear-gradient(135deg,#065f46,#10b981);color:#fff;text-align:center}
.hero h1{font-size:40px;margin:0 0 10px}.hero p{font-size:18px;margin:0;opacity:.92}
main{max-width:720px;margin:0 auto;padding:24px}a.row{display:flex;justify-content:space-between;align-items:center;background:#fff;border-radius:16px;padding:18px 22px;margin-bottom:12px;box-shadow:0 1px 3px rgba(0,0,0,.08);color:#1f2937;text-decoration:none;font-size:18px}
a.row code{font:600 17px ui-monospace,Menlo,monospace;color:#065f46}.get{background:#d1fae5;color:#065f46;border-radius:6px;padding:2px 8px;font:700 13px system-ui;margin-right:10px}</style></head>
<body><header>🍵 Cha Ghor</header><section class="hero"><h1>Cha Ghor API, in Go</h1><p>Compiled on your server with go build; JSON from the standard library.</p></section>
<main><a class="row" href="/api/menu"><span><span class="get">GET</span><code>/api/menu</code></span><span>The whole menu →</span></a>
<p style="color:#6b7280;text-align:center">POST /api/orders places an order</p></main></body></html>`))
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
