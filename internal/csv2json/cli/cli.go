package cli

import (
	"flag"
	"fmt"
)

type flags struct {
	csvFile string
}

func (f flags) CSVFile() string {
	return f.csvFile
}

func (f flags) String() string {
	return fmt.Sprintf("csvFile: %s", f.csvFile)
}

func ParseFlags() flags {
	csvFile := flag.String("csvfile", "", "path to CSV file with student data")

	flag.Parse()

	return flags{
		csvFile: *csvFile,
	}
}
