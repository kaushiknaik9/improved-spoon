package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaushiknaik9/improved-spoon/internal/config"
	"github.com/kaushiknaik9/improved-spoon/internal/handler"
	"github.com/kaushiknaik9/improved-spoon/internal/http/router"
	"github.com/kaushiknaik9/improved-spoon/internal/repository/postgres"
	"github.com/kaushiknaik9/improved-spoon/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db, err := postgres.NewDB(cfg.Storage.Url)
	if err != nil {
		log.Fatalf("can't connect to DB: %v", err)
	}
	defer db.Close()

	repo := postgres.NewURLRepository(db)

	urlService := service.NewURLService(repo)

	urlHandler := handler.NewURLHandler(urlService)

	r := chi.NewMux()

	urlRouter := router.NewURLHandler(urlHandler)

	urlRouter.RegisterURLRoutes(r)

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
