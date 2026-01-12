package main

import (
	"docs-site/internal/deps"
	"docs-site/internal/server/http/handlers"
	"docs-site/internal/server/http/routes"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const addr = ":8080"

func main() {
	d := deps.New()
	h := handlers.New(d)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	routes.Register(r, h)

	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
