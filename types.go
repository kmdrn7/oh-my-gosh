package main

type TemplateData struct {
	TotalRequest int
	Requests []RequestResult
}

type RequestResult struct {
	Command      string `json:"command"`
	ResultHeader string `json:"result_header"`
	ResultBody   string `json:"result_body"`
}
