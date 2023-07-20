package regex

import "regexp"

const (
	TOC             = "<ac:structured-macro ac:name=\"toc\"(.*?)/>"
	TOC_REPLACE     = ""
	INFO            = "<ac:structured-macro ac:name=\"info\" ac:schema-version=\"\\d\" ac:macro-id=\"[a-zA-Z0-9_.-]*\"><ac:rich-text-body><p>"
	INFO_REPLACE    = "<info-box>"
	INFOEND         = "</p></ac:rich-text-body></ac:structured-macro>"
	INFOEND_REPLACE = "</info-box>"
	DATE            = "<time datetime=\"([0-9_-]+)\" />"
	CONFLINK        = ""
)

func Regex(content string) (formatted string) {
	formatted = regexDate(content)
	formatted = regexInfo(formatted)
	formatted = regexTOC(formatted)
	return formatted
}

func regexTOC(content string) (formatted string) {
	reg := regexp.MustCompile(TOC)
	return reg.ReplaceAllString(content, TOC_REPLACE)
}

func regexInfo(content string) (formatted string) {
	reg := regexp.MustCompile(INFO)
	formatted = reg.ReplaceAllString(content, INFO_REPLACE)
	reg = regexp.MustCompile(INFOEND)
	return reg.ReplaceAllString(formatted, INFOEND_REPLACE)
}

func regexDate(content string) (formatted string) {
	//regexp.Compile()
	reg := regexp.MustCompile(DATE)
	if reg.Match([]byte(content)) {
		return reg.ReplaceAllString(content, reg.FindStringSubmatch(content)[1])
	}
	return content
}
