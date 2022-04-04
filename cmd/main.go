package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println()
	analyse()
}

func analyse() {
	r := newCSVReader()

	ds := Dataset{
		Headers: readHeaders(r),
	}

	ds.PrintHeaders()
}

func newCSVReader() *csv.Reader {
	file := getFilePath()
	r, e := newRowReader(file)

	if e != nil {
		panic(e)
	}

	return r
}

func getFilePath() string {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: path to covid data file")
		panic("Missing argument: path to covid data file")
	}
	return os.Args[1]
}

func newRowReader(file string) (*csv.Reader, error) {
	f, e := os.Open(file)
	if e != nil {
		return nil, e
	}
	return csv.NewReader(f), nil
}

func readHeaders(r *csv.Reader) []string {
	headers, e := r.Read()
	if e != nil {
		panic(e)
	}
	return headers
}

func printRows(r csv.Reader) {
	var e error

	for row, e := r.Read(); e == nil; row, e = r.Read() {
		fmt.Println(row)
	}

	if e != nil && e != io.EOF {
		fmt.Println("Expected EOF: ", e)
		panic(e)
	}
}

type Dataset struct {
	Headers []string
	Data    [][]string
}

func (ds *Dataset) PrintHeaders() {
	for _, h := range ds.Headers {
		fmt.Println(h)
	}
}
