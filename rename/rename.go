package rename

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Reaname() {
    nFlag := flag.Bool("n", false, "-n: no overwrite existing file")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	source := args[0]
	destination := args[1]

	sInfo, err := os.Lstat(source)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	if sInfo.IsDir() == true {
		fmt.Println("currently support only files")
	}

	dInfo, err := os.Lstat(source)

	if err == nil {
		if *nFlag == true {
			fmt.Println("error: file ", destination, " exists")

			os.Exit(1)
		}

		if dInfo.IsDir() == true {
			sourceName := filepath.Base(source)
			destination = destination + "/" + sourceName
		}
	}
	err = os.Rename(source, destination)

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(0)
}
