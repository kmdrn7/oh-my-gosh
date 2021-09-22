package main

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

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

	for _, curl := range curls {
		cmd := exec.Command("sh", "-c", curl)
		fmt.Println(cmd)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
			return
		} else {
			templateData.Requests = append(templateData.Requests, RequestResult{
				Command:      cmd.String(),
				ResultHeader: "-",
				ResultBody: string(output),
			})
		}
	}

	buf := &bytes.Buffer{}
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).Funcs(template.FuncMap{"unescape": unescape}).Parse(markdownTemplate))
	if err := tmpl.Execute(buf, *templateData); err != nil {
		fmt.Println("Failed to render template, this is a bug: ", err)
	}

	fmt.Println(buf.String())
}
