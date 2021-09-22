package main

var markdownTemplate = `### Results [`+"`Curl`"+`]
##
{{ range $idx, $val := .Requests }}
#### ********** {{ add1 $idx }} **********
*Command:*
`+"```sh\n"+
`{{unescape $val.Command}}`+
"\n```\n"+
`*Result Header:*`+
"\n```txt\n"+
`{{unescape $val.ResultHeader}}`+
"\n```\n"+
`*Result Body:*`+
"\n```json\n"+
`{{unescape $val.ResultBody}}`+
"\n```\n"+
"{{end}}"