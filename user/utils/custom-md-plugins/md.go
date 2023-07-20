package custom_md

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

func custom_md() (converter *md.Converter) {
	converter = md.NewConverter("", true, nil)
	converter.AddRules(InfoBoxRule())
	converter.Use(plugin.GitHubFlavored())
	converter.Use(plugin.ConfluenceCodeBlock())
	converter.AddRules(ConfLinkRule())
	return converter
}

func Format(content string) (formatted string) {
	converter := custom_md()
	formatted, _ = converter.ConvertString(content)
	return formatted
}
