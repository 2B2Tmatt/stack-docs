package handlers

import (
	"docs-site/web/templates/pages"
	"net/http"
)

func (h *Handler) Templ(w http.ResponseWriter, r *http.Request) {
	err := pages.Templ().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "issue rendering page", http.StatusInternalServerError)
	}
}
