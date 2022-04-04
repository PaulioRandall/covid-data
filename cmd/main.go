package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/covid-data/pkg/dataset"
	"github.com/PaulioRandall/covid-data/pkg/load"
)

func main() {
	fmt.Println()

	ds := loadDataset()

	fmt.Println(ds.Headers)

	// Next: Put the data into a SQLite database
}

func dataFile() string {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: path to covid data file")
		panic("Missing argument: path to covid data file")
	}
	return os.Args[1]
}

func loadDataset() dataset.Dataset {

	file := dataFile()
	headers, data, e := load.FromFile(file)

	if e != nil {
		panic(e)
	}

	return dataset.Dataset{
		Headers: headers,
		Data:    data,
	}
}
