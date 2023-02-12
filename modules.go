package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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

// Create files if not exist
func initFiles(files []string) {
	for i := 0; i < len(files); i++ {
		file := files[i]

		if _, err := os.Stat(file); err != nil {
			_, e := os.Create(file)

			if e != nil {
				fmt.Println(fmt.Sprintf(`Failed creating file: %v`, file))
			}
		}
	}

}

// The clothing size is 420, we resize it to 512 without losing quality
func resizeTemplate(link string) string {
	return fmt.Sprintf("https://tr.rbxcdn.com/%v/512/512/Image/Png", strings.Split(strings.Split(link, "https://tr.rbxcdn.com/")[1], "/")[0])
}
