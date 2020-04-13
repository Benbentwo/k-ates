package main

import (
	"github.com/Benbentwo/k-ates/kubectl"
	"github.com/Benbentwo/k-ates/parsers"
	"github.com/Benbentwo/k-ates/templates"
	"github.com/Benbentwo/k-ates/util"
	"net/http"
	"os"
	"strings"
)

// TODO's:
//  - it'd be nice to have a little console at the bottom that allows logs to be printed to a 'console' on the web page.
//  - Breadcrumbs would be nice to have a Type: (Auto|Manual|Off) flag where:
//     - Auto would split the url of the page into breadcrumbs
//     - Manual would allow manual addition of breadcrumbs
//     - Off just removes the sect.
//  - Make Html templates actual HTML files and use more modular parts.

var (
	Log   = util.Log
	DEBUG = util.DEBUG
	INFO  = util.INFO
	WARN  = util.WARN
	ERROR = util.ERROR
)

func handler(w http.ResponseWriter, r *http.Request) {
	var tmplt templates.Template
	rootPath := util.GetRootPath()
	filepath := r.URL.Path[len(rootPath):]
	Log(DEBUG, util.ColorInfo("FilePath: \"")+filepath+util.ColorInfo("\""))
	// log.Logger().Infof("FilePath: `%s`", util.ColorInfo(filepath))
	Log(DEBUG, util.ColorInfo("RootPath: \"")+rootPath+util.ColorInfo("\""))

	if strings.TrimSpace(filepath) == "" {
		Log(DEBUG, "Display Home Page")
		tmplt = templates.HOME
		util.LoadTemplate(w, tmplt, templates.Home{
			Filename: "Home Page",
			Headers: []templates.ButtonLinks{
				{
					Text: "Kubectl",
					Href: "/kubectl",
				},
			},
		})
	}

}
func loadCSVFileData(w http.ResponseWriter, r *http.Request, filepath string) (parsers.Table, error) {
	ext := filepath[strings.LastIndex(filepath, ".")+1:]
	var parser parsers.Parser
	switch ext {
	case "csv":
		parser = parsers.CSVParser{}
	case "prn":
		var columns []string
		columnStr := r.URL.Query().Get("columns")
		if strings.Trim(columnStr, " ") != "" {
			columns = strings.Split(columnStr, ",")
		}
		parser = parsers.FixedWidthParser{Columns: columns}
	default:
		util.HandleError(w, util.UnsupportedError{})
		return parsers.Table{}, util.UnsupportedError{}
	}
	reader, e := parsers.DecodeFile(filepath)
	if e != nil {
		util.HandleError(w, e)
		return parsers.Table{}, util.UnsupportedError{}
	}
	table, e := parser.Parse(reader)
	if e != nil {
		util.HandleError(w, e)
		return parsers.Table{}, util.UnsupportedError{}
	}
	table.Filename = filepath
	return table, nil
}

func main() {
	root := util.GetRootPath()
	err := templates.InitTemplates()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc(root, handler)
	http.HandleFunc(root+"kubectl/", kubectl.Handler)
	http.HandleFunc(root+"kubectl/get/pods/", kubectl.GetPodsHandler)
	_ = http.ListenAndServe(":"+port, nil)
}
