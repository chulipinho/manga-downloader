package main

import (
	"github.com/chulipinho/manga-dowloader/connectors"
	"github.com/chulipinho/manga-dowloader/manga"
)

func main() {
	c := connectors.Manga4lifeConnector{}
	manga.NewManga(c, "https://manga4life.com/read-online/Berserk-chapter-1.html")
}
