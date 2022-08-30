package util

import (
	"strings"
	"text/template"
)

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

var FuncMap = template.FuncMap{
	"replace": replace,
}
