package symlink

import (
	"fmt"
	"os"
	"path/filepath"
)

func Symlink() {
	if len(os.Args) == 1 {
		fmt.Println("provide a path to file, to test")
		os.Exit(1)
	}

	filePath := os.Args[1]

	fileinfo, err := os.Lstat(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		fmt.Printf("File: %s is symlink. \n", filePath)

		truePath, err := filepath.EvalSymlinks(filePath)

		if err != nil {
			fmt.Println("error during evaluating symlink: ", err)
			os.Exit(1)
		}

		fmt.Println("File path: ", truePath)
	}
}
