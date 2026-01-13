package handlers

import (
	"docs-site/web/templates/pages"
	"net/http"
)

func (h *Handler) Intro(w http.ResponseWriter, r *http.Request) {
	err := pages.Intro().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "issue rendering page", http.StatusInternalServerError)
	}
}
