package cli

import (
	"flag"
	"fmt"
)

type flags struct {
	csvFile  string
	jsonFile string
}

func (f flags) CSVFile() string {
	return f.csvFile
}

func (f flags) JSONFile() string {
	return f.jsonFile
}

func (f flags) String() string {
	return fmt.Sprintf("csvFile: %s, jsonFile: %s", f.csvFile, f.jsonFile)
}

func ParseFlags() flags {
	csvFile := flag.String("csvfile", "", "path to CSV file")
	jsonFile := flag.String("jsonfile", "", "path to JSON file")

	flag.Parse()

	return flags{
		csvFile:  *csvFile,
		jsonFile: *jsonFile,
	}
}
