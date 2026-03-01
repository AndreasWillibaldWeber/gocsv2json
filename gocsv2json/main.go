package main

import (
	"fmt"
	"os"

	"github.com/andreaswillibaldweber/gocsv2json/internal/csv2json"
	"github.com/andreaswillibaldweber/gocsv2json/internal/csv2json/cli"
)

func main() {
	flags := cli.ParseFlags()
	fmt.Printf("Flags>> %s \n\n", flags)

	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin stat error:", err)
		os.Exit(2)
	}
	if fi.Mode()&os.ModeCharDevice != 0 && flags.CSVFile() == "" {
		fmt.Fprintln(os.Stderr, "No stdin pipe detected. Usage: cat file.csv | csv2json")
		os.Exit(2)
	}

	input := os.Stdin
	if flags.CSVFile() != "" && fi.Mode()&os.ModeCharDevice != 0 {
		r, err := csv2json.CreateReader(flags.CSVFile())
		if err != nil {
			fmt.Fprintln(os.Stderr, "create reader error:", err)
			os.Exit(1)
		}
		input = r
	}
	defer input.Close()

	table, err := csv2json.NewCSVFrom(input, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "csv error:", err)
		os.Exit(1)
	}

	out := os.Stdout
	if flags.JSONFile() != "" {
		w, err := csv2json.CreateWriter(flags.JSONFile())
		if err != nil {
			fmt.Fprintln(os.Stderr, "create writer error:", err)
			os.Exit(1)
		}
		out = w
	}
	defer out.Close()

	json := csv2json.CSVtoJSON(table)
	err = csv2json.JSONTo(out, *json)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json error:", err)
		os.Exit(1)
	}
}
