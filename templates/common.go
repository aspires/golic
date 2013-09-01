package templates

import (
	"strings"
	"text/template"
)

const CopyrightTmpl = `{{define "COPYRIGHT"}}{{.Copyright}}{{if .Email}} <{{.Email}}>{{end}}{{if .URL}}, {{.URL}}{{end}}{{end}}`

type License struct {
	Name     string
	URL      string
	Template string
}

func List() []string {
	licenses := []string{}

	for i := range Licenses {
		licenses = append(licenses, Licenses[i].Name)
	}

	return licenses
}

func Load(name string) (*License, bool) {
	for i := range Licenses {
		if Licenses[i].Name == name {
			return &Licenses[i], true
		}
	}

	return nil, false
}

func Template(licTmpl string) (*template.Template, error) {
	tmpl := template.New("License")
	tmpl, err := tmpl.Parse(CopyrightTmpl)
	if err != nil {
		return nil, err
	}

	tmpl, err = tmpl.Parse(strings.TrimPrefix(licTmpl, "\n"))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
