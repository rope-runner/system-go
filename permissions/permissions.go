package permissions

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no file provided")
		os.Exit(1)
	}

	allSuccessfull := true

	for _, fPath := range os.Args[1:] {
		err := printPermissions(fPath)

		if err != nil {
			allSuccessfull = false
		}
	}

	if allSuccessfull == false {
		os.Exit(1)
	}

	os.Exit(0)
}

func printPermissions(fPath string) error {
	filestats, err := os.Stat(fPath)

	if err != nil {
		fmt.Println(fPath, ": ", strings.Split(err.Error(), "stat")[1])

		return err
	}

	fmt.Println(fPath, ": ", filestats.Mode())

	return nil
}


