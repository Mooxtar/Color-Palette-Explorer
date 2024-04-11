package main

import (
	"github.com/go-chi/chi"
	"github.com/Mooxtar/color-palette-explorer/internal/config"
	"github.com/Mooxtar/color-palette-explorer/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/most", handlers.Repo.Most)
	mux.Get("/generate", handlers.Repo.Generate)

	return mux
}
