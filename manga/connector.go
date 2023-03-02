package manga

import "io"

type connector interface {
	GetParametersFromBody(io.ReadCloser) (title, path, currentChapter string)
}
