package main

import (
	"fmt"
	"github.com/Benbentwo/k-ates/parsers"
	"github.com/Benbentwo/k-ates/templates"
	"github.com/Benbentwo/k-ates/util"
	"github.com/fatih/color"
	"github.com/jenkins-x/jx/pkg/log"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	white  = color.New(color.FgWhite).SprintFunc()
)

type UnsupportedError struct{}

func (e UnsupportedError) Error() string {
	return "Unsupported file type"
}

func handler(w http.ResponseWriter, r *http.Request) {
	var tmplt templates.Template
	rootPath := os.Getenv("ROOT")
	if rootPath == "" {
		rootPath = "/"
	}
	filepath := r.URL.Path[len(rootPath):]
	log.Logger().Infof("FilePath: `%s`", util.ColorInfo(filepath))

	if strings.TrimSpace(rootPath) == "" {
		tmplt = templates.HOME
		loadTemplate(w, tmplt, parsers.Table{
			Filename: "Home Page",
		})
	} else {
		tmplt = templates.TABLE
		table, err := loadCSVFileData(w, r, filepath)
		if err != nil {
			log.Logger().Errorf("%s: %s", red("ERRORs"), white(err))
		}
		loadTemplate(w, tmplt, table)
	}

	// t, _ := template.New("table").Parse(templates.Table)
	// t.Execute(w, table)
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
		handleError(w, UnsupportedError{})
		return parsers.Table{}, UnsupportedError{}
	}
	reader, e := parsers.DecodeFile(filepath)
	if e != nil {
		handleError(w, e)
		return parsers.Table{}, UnsupportedError{}
	}
	table, e := parser.Parse(reader)
	if e != nil {
		handleError(w, e)
		return parsers.Table{}, UnsupportedError{}
	}
	table.Filename = filepath
	return table, nil
}
func loadTemplate(w http.ResponseWriter, temp templates.Template, content interface{}) {
	t, _ := template.New("table").Parse(temp.GetTemplate())
	err := t.Execute(w, content)
	if err != nil {
		handleError(w, err)
		return
	}
}
func handleError(w http.ResponseWriter, e error) {
	err := e.Error()
	var status int
	switch t := e.(type) {
	case *os.PathError:
		err = fmt.Sprintf("File '%s' not found", t.Path)
		status = 404
	case UnsupportedError:
		status = 400
	case parsers.EmptyFileError:
		status = 400
	case parsers.ColNotFound:
		status = 400
	default:
		err = fmt.Sprintf("Something went wrong that was unforseen: %s", e.Error())
		status = 500
	}
	http.Error(w, err, status)
}

func main() {
	root := os.Getenv("ROOT")
	if root == "" { // These aren't needed on the dockerized version as the ARG default sets them.
		root = "/" //    They are still nice to have for local go building though.
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc(root, handler)
	_ = http.ListenAndServe(":"+port, nil)
}
