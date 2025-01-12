package which

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const aflagName = "a"


func Which() {
    helpFlag := flag.Bool("help", false, "-help: print the docs of the command")

	aFlag := flag.Bool(aflagName, false, "-a: print all found entries for provided arguments")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		os.Exit(1)
	}

	sysPath := os.Getenv("PATH")
	allFound := true

	for _, command := range args {
		allFound = checkCommand(command, sysPath, *aFlag)
	}

	if allFound == true {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func checkCommand(command string, sysPath string, aFlag bool) bool {
	sysPathParts := strings.Split(sysPath, ":")
	found := false

	for _, direcotry := range sysPathParts {
		commandPath := direcotry + "/" + command

		cmdFile, err := os.Lstat(commandPath)

		if err == nil {
			if cmdFile.Mode()&0o0111 != 0 {
				found = true

				fmt.Println(commandPath)

				if aFlag == false {
					break
				}
			}
		}
	}

	if found == false {
		fmt.Println("command not found: ", command)
	}

	return found
}

