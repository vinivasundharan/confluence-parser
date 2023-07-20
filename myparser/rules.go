package myparser

import (
	"bufio"
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func InfoBoxRule() (infoboxRule md.Rule) {
	infoBoxRule := md.Rule{

		Filter: []string{"info-box"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			breakCharacter := "<br>"
			text := content
			scanner := bufio.NewScanner(strings.NewReader(text))
			infoboxmd := "> &#9432; INFOBOX" + breakCharacter + "\n"
			for scanner.Scan() {
				if scanner.Text() != "" {
					infoboxmd = infoboxmd + scanner.Text() + breakCharacter + "\n"
				}
			}
			md := fmt.Sprintf("%s", infoboxmd)
			return &md
		},
	}
	return infoBoxRule
}
