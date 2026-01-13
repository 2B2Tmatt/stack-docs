package deps

import (
	"io/fs"
	"net/http"
)

type Deps struct {
	AssetsFS   http.FileSystem
	SnippetsFS fs.FS
}

func New(snippets fs.FS) *Deps {
	return &Deps{
		AssetsFS:   http.Dir("web/static"),
		SnippetsFS: snippets,
	}
}
