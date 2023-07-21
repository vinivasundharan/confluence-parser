package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"markdown/markdown/utils/confluence-connector"
	custom_md "markdown/markdown/utils/custom-md-plugins"
	"markdown/markdown/utils/regex"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func download(w http.ResponseWriter, r *http.Request) {
	// Set headers for the file download.
	w.Header().Set("Content-Disposition", "attachment; filename=response.md")
	w.Header().Set("Content-Type", "application/octet-stream")

	// Open and read the file content.
	fileContent, err := ioutil.ReadFile("response.md")
	if err != nil {
		http.Error(w, "Failed to read the file", http.StatusInternalServerError)
		return
	}

	// Write the file content to the HTTP response.
	_, err = w.Write(fileContent)
	if err != nil {
		http.Error(w, "Failed to write the file content to the response", http.StatusInternalServerError)
		return
	}
}

type PageData struct {
	Success  bool
	FileName string
}

func main() {
	os.Remove("response.md")
	//var confContent confluence.Content
	http.HandleFunc("/conf2md", conf2md)
	http.HandleFunc("/download", download)
	http.ListenAndServe(":8080", nil)
	//var content Content 65883
}

func generateMD(contentID string) (filename string) {
	fmt.Println("in generateMD")
	os.Remove("response.md")
	var confContent confluence.Content
	confContent.ID = contentID
	bodyrequest := confluence.ConfPageContent(confContent.ID)
	json.Unmarshal(bodyrequest, &confContent)
	title := "# " + confContent.Title + "\n"
	formatted_html := regex.Regex(confContent.Body.Storage.Value)
	markdown := custom_md.Format(formatted_html)
	f, _ := os.OpenFile("response.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(title)
	f.WriteString(markdown)
	return "response.md"
}

// Renderer function
func conf2md(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("render_templates/forms.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		pageData := PageData{}
		pageData.FileName = generateMD(r.FormValue("email"))
		pageData.Success = true

		t, err := template.ParseFiles("render_templates/forms.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}
		t.Execute(w, pageData)
	} else if r.Method == "HEAD" || r.Method == "GET" {
		// Download request
		fmt.Print(r.Method)
		download(w, r)
	}

}
