package main

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"strings"
)

type CurlResponse struct {
	Command string `json:"command"`
	Output string `json:"output"`
}

func main() {
	Curlfile, err := os.ReadFile("Curlfile")
	if err != nil {
		fmt.Println("Error reading Curlfile:", err)
	}

	curls := strings.Split(string(Curlfile), "\n")

	requests := &[]RequestResult{}
	templateData := &TemplateData{
		TotalRequest: len(curls),
		Requests:     *requests,
	}

	resChan := make(chan RequestResult)
	for _, curl := range curls {
		go Curl(curl, resChan)
	}

	for range curls {
		templateData.Requests = append(templateData.Requests, <-resChan)
	}

	buf := &bytes.Buffer{}
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).Funcs(template.FuncMap{"unescape": unescape}).Parse(markdownTemplate))
	if err := tmpl.Execute(buf, *templateData); err != nil {
		fmt.Println("Failed to render template, this is a bug: ", err)
	}

	fmt.Println(buf.String())
}
