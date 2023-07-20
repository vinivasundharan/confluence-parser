package main

import (
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://blog.golang.org/godoc-documenting-go-code"

	doc, err := goquery.NewDocument(url)
	docf, _ := os.OpenFile("response.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	content := doc.Find("#content")
	docf.WriteString(content.Text())
	conv := md.NewConverter(md.DomainFromURL(url), true, nil)
	markdown := conv.Convert(content)
	f, _ := os.OpenFile("response.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(markdown)
	//fmt.Println(markdown)
}
