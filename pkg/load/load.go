package load

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func FromFile(file string) (headers []string, data [][]string, e error) {

	f, e := os.Open(file)
	if e != nil {
		return nil, nil, e
	}
	defer f.Close()

	r := csv.NewReader(f)

	headers, e = r.Read()
	if e != nil {
		return nil, nil, e
	}

	for row, e := r.Read(); e == nil; row, e = r.Read() {
		data = append(data, row)
	}

	if e != nil && e != io.EOF {
		fmt.Println("Expected EOF: ", e)
		return nil, nil, e
	}

	return headers, data, nil
}
