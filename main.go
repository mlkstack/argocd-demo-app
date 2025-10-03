// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	appName    = "demo-argocd-app"
	appVersion = "v1.0.0"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Application: %s\nVersion: %s\n", appName, appVersion)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(msg))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting %s on port %s (version %s)", appName, port, appVersion)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
