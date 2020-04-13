package util

import (
	"net/http"
	"os"
)

func GetRootPath() string {
	rootPath := os.Getenv("ROOT")
	if rootPath == "" {
		rootPath = "/"
	}
	return rootPath
}
func GetFilePath(r *http.Request) string {
	rootPath := GetRootPath()
	return r.URL.Path[len(rootPath):]
}
