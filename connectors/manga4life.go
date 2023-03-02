package connectors

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type Manga4lifeConnector struct{}

func (m Manga4lifeConnector) GetParametersFromBody(r io.ReadCloser) (title, path, currentChapter string) {
	var inMainFunc bool
	var filledVars int

	s := bufio.NewScanner(r)
	regex := regexp.MustCompile(`=\s.*?;`)

	for s.Scan() {
		if strings.Contains(s.Text(), "MainFunction") {
			inMainFunc = true
		}
		if !inMainFunc {
			continue
		}

		if strings.Contains(s.Text(), "CurPathName") {
			path = clearString(regex.FindString(s.Text()))
			filledVars++
		}
		if strings.Contains(s.Text(), "IndexName") {
			title = clearString(regex.FindString(s.Text()))
			filledVars++
		}
		if strings.Contains(s.Text(), "CurChapter") {
			currentChapter = getChapter(clearString(regex.FindString(s.Text())))
			filledVars++
		}

		if filledVars == 3 {
			break
		}
	}

	return
}

func getChapter(s string) string {
	regex := regexp.MustCompile(`"Chapter":.*?\,`)
	chap := regex.FindString(s)

	return chap[11 : len(chap)-2]
}

func clearString(s string) string {
	return s[3 : len(s)-2]
}
