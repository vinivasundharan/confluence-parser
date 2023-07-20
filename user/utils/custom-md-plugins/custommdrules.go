package custom_md

import (
	"bufio"
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func InfoBoxRule() (infoBoxRule md.Rule) {
	infoBoxRule = md.Rule{

		Filter: []string{"info-box"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			breakCharacter := "<br>"
			//text := content
			scanner := bufio.NewScanner(strings.NewReader(content))
			infoboxmd := "\n> &#9432; INFOBOX" + breakCharacter + "\n"
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

func ConfLinkRule() (confLinkRule md.Rule) {
	confLinkRule = md.Rule{
		Filter: []string{"ac:link"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			if selec.Nodes[0].FirstChild.Attr[0].Key == "ri:content-title" {
				attr := selec.Nodes[0].FirstChild.Attr
				md := fmt.Sprintf("[%s](%s)", selec.Nodes[0].LastChild.FirstChild.FirstChild.Data, attr[0].Val)
				return &md
			}
			if selec.Nodes[0].FirstChild.Attr[0].Key == "ri:space-key" {
				spaceName := selec.Nodes[0].FirstChild.Attr[0].Val
				contentName := selec.Nodes[0].FirstChild.Attr[1].Val
				md := fmt.Sprintf("[%s](%s/%s)", selec.Nodes[0].LastChild.FirstChild.FirstChild.Data, spaceName, contentName)
				return &md
			}
			return md.String("")
		},
	}
	return confLinkRule
}
