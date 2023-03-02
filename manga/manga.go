package manga

import (
	"net/http"
)

func NewManga(c connector, url string) manga {
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	title, path, currentChapter := c.GetParametersFromBody(r.Body)

	return manga{
		Title:          title,
		Path:           path,
		CurrentChapter: currentChapter,
	}
}

type manga struct {
	Title          string
	Path           string
	CurrentChapter string
}
