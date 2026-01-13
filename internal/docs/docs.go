package docs

import "io/fs"

type Doc struct {
	Slug     string
	Title    string
	Sections []Section
}

type Section struct {
	Part string
}

func ReadSnippet(fsys fs.FS, path string) (string, error) {
	b, err := fs.ReadFile(fsys, path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
