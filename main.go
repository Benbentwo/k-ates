package main

import (
	"fmt"
	"github.com/Benbentwo/k-ates/parsers"
	"github.com/Benbentwo/k-ates/templates"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type UnsupportedError struct{}

func (e UnsupportedError) Error() string {
	return "Unsupported file type"
}

func handler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT")
	if rootPath == "" {
		rootPath = "/"
	}
	filepath := r.URL.Path[len(rootPath):]
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
		return
	}
	reader, e := parsers.DecodeFile(filepath)
	if e != nil {
		handleError(w, e)
		return
	}
	table, e := parser.Parse(reader)
	table.Filename = filepath
	if e != nil {
		handleError(w, e)
		return
	}
	t, _ := template.New("table").Parse(templates.Table)
	t.Execute(w, table)
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
	http.HandleFunc(root, handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	_ = http.ListenAndServe(":"+port, nil)
}
