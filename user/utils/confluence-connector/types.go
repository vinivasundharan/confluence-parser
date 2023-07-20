package confluence

type Content struct {
	Title      string `json:"title"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Body       Body   `json:"body"`
	Link       Link   `json:"_links"`
	Expandable struct {
		Space string `json:"space"`
	} `json:"_expandable"`
}

type Body struct {
	Storage struct {
		Value          string `json:"value"`
		Representation string `json:"representation"`
	} `json:"storage"`
}

type Link struct {
	WebUI string `json:"webui"`
	Self  string `json:"self"`
	//Key string `json:"key"`
}
