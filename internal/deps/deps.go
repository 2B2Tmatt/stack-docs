package deps

import "net/http"

type Deps struct {
	AssetsFS http.FileSystem
}

func New() *Deps {
	return &Deps{
		AssetsFS: http.Dir("web/static"),
	}
}
