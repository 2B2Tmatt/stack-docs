package routes

import (
	"docs-site/internal/server/http/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Register(r chi.Router, h *handlers.Handler) {
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(h.Deps.AssetsFS)))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home page")
	})
}
