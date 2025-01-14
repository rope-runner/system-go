package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadColumn(r io.Reader, coln int) {
	scanner := bufio.NewScanner(r)
	rows := 0

	for scanner.Scan() {
		cols := strings.Fields(scanner.Text())
		rows += 1

		if len(cols) <= coln-1 {
			fmt.Printf("Column %d is not present in row %d", coln, rows)
		} else {
			fmt.Println(cols[coln-1])
		}
	}
}

func main() {
	filename := os.Args[1]

	file, err := os.OpenFile(filename, os.O_RDONLY, 0o666)

	if err != nil {
		fmt.Println(err)

		return
	}

	defer file.Close()

	fmt.Println("===========")

	br := bufio.NewReader(file)

	br.Reset(file)

	for {
		sl, err := br.ReadString('\n')

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(string(sl))
	}
}
