package routes

import (
	"docs-site/internal/server/http/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Register(r chi.Router, h *handlers.Handler) {
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(h.Deps.AssetsFS)))
	r.Get("/docs/intro", h.Intro)
	r.Get("/docs/go", h.Go)
	r.Get("/docs/templ", h.Templ)
}
