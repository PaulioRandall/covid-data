package dataset

import (
	"fmt"
)

type Dataset struct {
	Headers []string
	Data    [][]string
}

func (ds *Dataset) PrintHeaders() {
	for _, h := range ds.Headers {
		fmt.Println(h)
	}
}
