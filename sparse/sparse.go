package sparse

import (
	"fmt"
	"os"
	"strconv"
)

func CreateSparseFile() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("not enough arguments")

		os.Exit(1)
	}

	filename := args[1]
	SIZE, err := strconv.ParseInt(args[2], 10, 64)

	if err != nil {
		fmt.Printf("size is not in proper format: %s \n", args[2])

		os.Exit(1)
	}

	_, err = os.Stat(filename)

	if err == nil {
		fmt.Println("file already exist")

		os.Exit(1)
	}

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	defer file.Close()

	_, err = file.Seek(SIZE-1, 0)

	if err != nil {
		fmt.Println(err)

		os.Remove(filename)
		file.Close()

		os.Exit(1)
	}

	_, err = file.Write([]byte{0})

	if err != nil {
		fmt.Println("error writing to file")

		os.Remove(filename)
		file.Close()

		os.Exit(1)
	}

}
