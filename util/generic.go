package util

import (
	"os"
)

func GetRootPath() string {
	rootPath := os.Getenv("ROOT")
	if rootPath == "" {
		rootPath = "/"
	}
	return rootPath
}
