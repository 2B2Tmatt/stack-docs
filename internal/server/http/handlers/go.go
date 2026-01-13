package handlers

import (
	"docs-site/internal/docs"
	"docs-site/web/templates/pages"
	"net/http"
)

func (h *Handler) Go(w http.ResponseWriter, r *http.Request) {
	initProj, _ := docs.ReadSnippet(h.Deps.SnippetsFS, "init_project.sh")
	simpleServer, _ := docs.ReadSnippet(h.Deps.SnippetsFS, "server_minimal.go.txt")
	runProgram, _ := docs.ReadSnippet(h.Deps.SnippetsFS, "run_program.sh")
	jsonStruct, _ := docs.ReadSnippet(h.Deps.SnippetsFS, "json_struct.go.txt")
	errorHandling, _ := docs.ReadSnippet(h.Deps.SnippetsFS, "error_handling.go.txt")

	err := pages.Go(initProj, simpleServer, runProgram, jsonStruct, errorHandling).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "issue rendering", http.StatusInternalServerError)
	}
}
