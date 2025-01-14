package getrows

import (
	"io"
	"strings"
)


func GetRows(r io.Reader, bufSize int, rowsLimit int) (error, []string) {
	buf := make([]byte, bufSize)
	rows := make([]string, 0, 100)
	partial := ""
	str := ""

	for {
		n, err := r.Read(buf)

		if err == io.EOF {
			return nil, rows
		}

		if err != nil {
			return err, nil
		}

		str = string(buf[:n])
		processingStr := str

		if len(partial) != 0 {
			processingStr = partial + str
		}

		if strings.Index(processingStr, "\n") != -1 {
			for {
				i := strings.Index(processingStr, "\n")

				if i == 0 {
					tStr := processingStr[1:]
					nextTerminator := strings.Index(tStr, "\n")

					if nextTerminator != -1 {
						processingStr = tStr
						rows = append(rows, processingStr[:nextTerminator])

						if len(rows) == rowsLimit {
							return nil, rows
						}

						processingStr = processingStr[nextTerminator:]
					} else {
						partial = processingStr[1:]
						break
					}
				} else if i == -1 {
					partial = processingStr
					break
				} else {
					rows = append(rows, processingStr[:i])

					if len(rows) == rowsLimit {
						return nil, rows
					}

					processingStr = processingStr[i:]
				}
			}
		} else {
			partial = processingStr
		}
	}
}
