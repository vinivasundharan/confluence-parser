package confluence_old

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// User defines user informations
type Confluence struct {
	EndPoint string `default:"https://vinivasundharan.atlassian.net/wiki/rest/api"`
	APIToken string `default:"ATATT3xFfGF0NPVT32-8i0-dONLMQTaqSCH-Lr7gC6Ew04iwzfdbC-MIB4cATjeGeCxGKm1_5KpsRwGU6ccVb4QiVthhLa4Gn7J55ht9ekInk1LJm0mogigIigVN-9tQrPg_t0RB6tU6gI9lKdISKXe1TMF5eNzTxt1QVv_gRDwgb6CFqvEj8N0=E60B440A"`
}

// GetStringInBetween returns empty string if no start or end string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}
func listelement(html1 string) (words []string) {
	fmt.Println(html1)
	html1 = strings.ReplaceAll(html1, "<br />", "\n")
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html1))
	fmt.Println(html1)
	//listelement, _ := html.Parse(strings.NewReader(html1))
	words = doc.Find("li").Map(func(i int, sel *goquery.Selection) string {
		fmt.Println(i)
		fmt.Println(doc.Nodes)
		fmt.Println()
		fmt.Print(sel)
		return fmt.Sprintf("\n%s", sel.Text())
	})

	fmt.Println(words)
	return words
}
func Getlist(str string) (result string) {
	start := "<ul>"
	end := "</ul>"
	list := GetStringInBetween(str, start, end)
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return list
}
func parse1(text string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))
	var vals []string
	var isLi bool
	for {
		tt := tkn.Next()
		//fmt.Printf("%+v\n", tt)
		switch {
		case tt == html.ErrorToken:
			return vals
		case tt == html.StartTagToken:
			fmt.Println("case StartTagToken")
			t := tkn.Token()
			isLi = t.Data == "li"

			fmt.Printf("%+v\t%+v\t%+v\n", t.Data, isLi, t.Type)
			if isLi {
				tkn.Next()
				tp := tkn.Token()
				fmt.Printf("%+v\t%+v\t%+v\n", tp.Data, isLi, tp.Type)
			}

		case tt == html.TextToken:
			fmt.Println("case TextToken")
			t := tkn.Token()
			fmt.Printf("%+v\t%+v\t%+v\n", t.Data, isLi, t.Type)
			if isLi {
				vals = append(vals, t.Data)
				tkn.Next()
				//tp := tkn.Token()
				//fmt.Printf("%+v", tp)

			}
			isLi = false
		case tt == html.TokenType(html.ElementNode):
			fmt.Println("case ElementNode")
			t := tkn.Token()
			fmt.Printf("%+v\t%+v\t%+v\n", t.Data, isLi, t.Type)
			if isLi {
				vals = append(vals, t.Data)

			}
			isLi = false
		}

	}
}
