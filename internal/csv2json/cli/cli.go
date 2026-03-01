package cli

import (
	"flag"
	"fmt"
)

type flags struct {
	noHeader bool
	csvFile  string
	jsonFile string
}

func (f flags) Header() bool {
	return !f.noHeader
}

func (f flags) CSVFile() string {
	return f.csvFile
}

func (f flags) JSONFile() string {
	return f.jsonFile
}

func (f flags) String() string {
	return fmt.Sprintf("noHeader: %t, csvFile: %s, jsonFile: %s", f.noHeader, f.csvFile, f.jsonFile)
}

func ParseFlags() flags {
	noHeader := flag.Bool("noheader", false, "has header row")
	csvFile := flag.String("csvfile", "", "path to CSV file")
	jsonFile := flag.String("jsonfile", "", "path to JSON file")

	flag.Parse()

	return flags{
		noHeader: *noHeader,
		csvFile:  *csvFile,
		jsonFile: *jsonFile,
	}
}
