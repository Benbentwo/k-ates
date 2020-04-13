package templates

import (
	"io/ioutil"
)

type Template string

var (
	HOME  Template
	TABLE Template
	K8    Template
)

const (
	RefreshText = "Refresh"
	RefreshUrl  = "."
)

func InitTemplates() error {
	tmpl, err := CreateTemplate("templates/home.html")
	if err != nil {
		return err
	}
	HOME = Template(tmpl)

	tmpl, err = CreateTemplate("templates/table.html")
	if err != nil {
		return err
	}
	TABLE = Template(tmpl)

	tmpl, err = CreateTemplate("templates/k8-table.html")
	if err != nil {
		return err
	}
	K8 = Template(tmpl)
	return nil
}

func (template Template) GetTemplate() string {
	return string(template)

}

func CreateTemplate(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

type ButtonLinks struct {
	Text string
	Href string
}
type Home struct {
	Filename string
	Headers  []ButtonLinks
}

type K8Template struct {
	Filename    string
	Buttons     []ButtonLinks
	BreadCrumbs []ButtonLinks
	Headers     []string
	Rows        [][]string
}
type TableTemplate struct {
	Filename string
	Headers  []string
	Rows     [][]string
}

var RefreshButton = ButtonLinks{
	Href: RefreshUrl,
	Text: RefreshText,
}
