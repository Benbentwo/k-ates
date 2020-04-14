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

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
