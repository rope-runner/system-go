package remove

import (
	"flag"
	"fmt"
	"os"
)

func Remove() {
	rFlag := flag.Bool("r", false, "-r: remove content of directory recursively")
	fFlag := flag.Bool("f", false, "-f: ingore non existent files and directories, never prompt")
	vFlag := flag.Bool("v", false, "-v: verbose, log actions")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("no file(s) provided")
		os.Exit(1)
	}

	for _, fPath := range flag.Args() {
		filestat, err := os.Lstat(fPath)

		if err != nil {
			if *fFlag == false {
				fmt.Println(err)
			}
		}

		if filestat.IsDir() == true {
			if *rFlag == false {
				if *fFlag == false {
					fmt.Println("error: ", fPath, " is directory")
					continue
				}
			} else {
				removeDirContent(fPath, *fFlag, *vFlag)
				removeFile(fPath, *fFlag, *vFlag)
			}
		} else {
			removeFile(fPath, *fFlag, *vFlag)
		}

	}
}

func removeFile(fPath string, fFlag bool, vFlag bool) {
	err := os.Remove(fPath)

	if err != nil {
		if fFlag == false {
			fmt.Println("error: removing ", fPath, " ", err)
		}
	} else {
		if vFlag == true {
			fmt.Println("removed: ", fPath)
		}
	}
}

func removeDirContent(dirPath string, fFlag bool, vFlag bool) {
	dirinfo, err := os.ReadDir(dirPath)

	if err != nil {
		if fFlag == false {
			fmt.Println("error: reading ", dirPath, " ", err)
		}

		return
	}

	for _, entry := range dirinfo {
		entryPath := dirPath + "/" + entry.Name()

		entryStat, err := os.Lstat(entryPath)

		if err != nil {
			if fFlag == false {
				fmt.Println(err)
			}

			continue
		}

		if entryStat.IsDir() == true {
			removeDirContent(entryPath, fFlag, vFlag)

			removeFile(entryPath, fFlag, vFlag)
		} else {
			removeFile(entryPath, fFlag, vFlag)
		}
	}
}
