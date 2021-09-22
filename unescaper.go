package main

import "html/template"

func unescape(s string) template.HTML {
	return template.HTML(s)
}
