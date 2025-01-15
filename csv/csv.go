package csv

import (
	"encoding/csv"
	"fmt"
	"io"
)

func ReadFromCSV(file io.Reader, comma rune) error {
	csvR := csv.NewReader(file)

	csvR.Comma = ';'

	for {
		record, err := csvR.Read()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		for i, v := range record {
			fmt.Printf("%d: %s \n", i, v)
		}

		fmt.Println()
	}
}

func WriteToCSV(file io.WriteCloser) error {
	testData := [][]string{{"username", "email"}, {"rkrkrk", "rr@mail.com"}}

	csvW := csv.NewWriter(file)

	for _, record := range testData {
		err := csvW.Write(record)

		if err != nil {
			file.Close()

			return err
		}
	}

	csvW.Flush()

	return nil
}
