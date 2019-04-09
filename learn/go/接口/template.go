package main

import (
	"html/template"
	"os"
	"github.com/adonovan/gopl.io/ch4/github"
)

const templ  = `{{ .TotalCount }} issues
number: {{.number}}
user:
title:
age:`

type database map[string]string

func main() {
	report, _ := template.New("report").Parse(templ)
	report.Execute(os.Stdout, database{"TotalCount": "32", "number": "2"})
}
