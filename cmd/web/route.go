package main

import (
	"net/http"

	"github.com/catehulu/rigged-coin/internal/config"
	"github.com/catehulu/rigged-coin/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/boards", handlers.Repo.GetBoards)
	mux.Post("/boards", handlers.Repo.PostBoards)

	return mux
}
