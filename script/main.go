package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	err := listMyDirs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listMyDirs() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return filepath.Walk(pwd, filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("Dir: %s\n\tMod Time: %s\n", path, info.ModTime().Format(time.RFC3339))
			return nil
		}
		fmt.Printf("File: %s\n", path)
		return nil
	}))
}
