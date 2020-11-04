package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("all go files: ")

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			if filepath.Ext(filepath.Base(path)) == ".go" || filepath.Ext(filepath.Base(path)) == ".cgo" {
				fmt.Printf("%q\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
	}
}
