package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaushiknaik9/improved-spoon/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	r := chi.NewMux()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]any{
			"success": "true",
			"message": "Spoon running at 3000",
		})
	})

	http.ListenAndServe(cfg.Host, r)
}
