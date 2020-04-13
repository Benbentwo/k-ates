package templates

var (
	HOME  Template = Template(homePage)
	TABLE Template = Template(Table)
	K8    Template = Template(K8Table)
)

func (template Template) GetTemplate() string {
	return string(template)

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
