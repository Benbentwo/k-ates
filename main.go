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

var (
	Log   = util.Log
	DEBUG = util.DEBUG
	INFO  = util.INFO
	WARN  = util.WARN
	ERROR = util.ERROR
)

func handler(w http.ResponseWriter, r *http.Request) {
	util.Log(util.INFO, "handler")
	rootPath := util.GetRootPath()
	util.LoadTemplate(w, templates.HOME, templates.Home{
		Filename: "Home Page",
		Headers: []templates.ButtonLinks{
			{
				Text: "Kubectl",
				Href: rootPath + "kubectl",
			},
		},
	})

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
	util.Log(util.DEBUG, "init templates")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	kubectl.NewKubectl()

	http.HandleFunc(root, handler)
	util.Log(util.DEBUG, "root")
	http.HandleFunc(root+"kubectl/", kubectl.Handler)
	util.Log(util.DEBUG, "Kubectl")
	http.HandleFunc(root+"kubectl/get/pods/", kubectl.GetPodsHandler)
	util.Log(util.DEBUG, "Kubectl Pods")
	_ = http.ListenAndServe(":"+port, nil)
}
