package util

import (
	"fmt"
	"github.com/Benbentwo/k-ates/parsers"
	"github.com/Benbentwo/k-ates/templates"
	"html/template"
	"net/http"
	"os"
)

type UnsupportedError struct{}

func (e UnsupportedError) Error() string {
	return "Unsupported file type"
}

func LoadTemplate(w http.ResponseWriter, temp templates.Template, content interface{}) {
	t, _ := template.New("table").Parse(temp.GetTemplate())
	err := t.Execute(w, content)
	if err != nil {
		HandleError(w, err)
		return
	}
}

func HandleError(w http.ResponseWriter, e error) {
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
