package main

import (
	"errors"
	"os"
)

// Create directories if not exist
func initDirs(paths []string) {
	for i := 0; i < len(paths); i++ {
		path := paths[i]

		if _, err := os.Stat(path); err == nil {
		} else if errors.Is(err, os.ErrNotExist) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			os.MkdirAll(path, os.ModePerm)
		}
	}
}
