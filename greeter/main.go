// Command greeter runs a small, stateless HTTP service exposing GET /hello.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "9090"
	defaultName = "World"
)

// Greeting matches the Greeting schema in openapi.yaml.
type Greeting struct {
	Message string `json:"message"`
}

// errorResponse matches the organization's JSON error-body convention.
type errorResponse struct {
	Error string `json:"error"`
}

type config struct {
	port        string
	defaultName string
}

func loadConfig() config {
	cfg := config{port: defaultPort, defaultName: defaultName}
	if v := os.Getenv("PORT"); v != "" {
		cfg.port = v
	}
	if v := os.Getenv("GREETER_DEFAULT_NAME"); v != "" {
		cfg.defaultName = v
	}
	return cfg
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("greeter: failed to encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func helloHandler(cfg config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, fmt.Sprintf("method %s not allowed", r.Method))
			return
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			name = cfg.defaultName
		}

		writeJSON(w, http.StatusOK, Greeting{Message: fmt.Sprintf("Hello, %s!", name)})
	}
}

func newMux(cfg config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler(cfg))
	return mux
}

func main() {
	cfg := loadConfig()
	addr := ":" + cfg.port

	log.Printf("greeter: listening on %s", addr)
	if err := http.ListenAndServe(addr, newMux(cfg)); err != nil {
		log.Fatalf("greeter: server failed: %v", err)
	}
}
