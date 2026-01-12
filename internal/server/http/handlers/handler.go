package handlers

import "docs-site/internal/deps"

type Handler struct {
	Deps *deps.Deps
}

func New(d *deps.Deps) *Handler {
	return &Handler{Deps: d}
}
