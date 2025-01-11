package pwd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const pflagName = "P"

func Pwd() {
	pFlag := flag.Bool(pflagName, false, "-P: resolve symlinks")

	flag.Parse()

	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pwd)
	}

	if *pFlag == false {
		return
	}

	realpath, err := filepath.EvalSymlinks(pwd)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(realpath)
	}
}
