package main

import (
	"embed"
	"io/fs"
)

//go:embed snippets/**
var snippetsEmbedFS embed.FS

func SnippetsFS() fs.FS {
	sub, err := fs.Sub(snippetsEmbedFS, "snippets")
	if err != nil {
		panic(err)
	}
	return sub
}
