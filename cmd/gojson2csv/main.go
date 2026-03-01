package main

import (
	"fmt"
	"os"

	"github.com/andreaswillibaldweber/gocsv2json/internal/cli"
	"github.com/andreaswillibaldweber/gocsv2json/internal/iio"
	"github.com/andreaswillibaldweber/gocsv2json/internal/json2csv"
)

func main() {
	flags := cli.ParseFlags()
	fmt.Printf("Flags>> %s \n\n", flags)

	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin stat error:", err)
		os.Exit(2)
	}
	if fi.Mode()&os.ModeCharDevice != 0 && flags.JSONFile() == "" {
		fmt.Fprintln(os.Stderr, "No stdin pipe detected. Usage: cat file.json | json2csv")
		os.Exit(2)
	}

	input := os.Stdin
	if flags.JSONFile() != "" && fi.Mode()&os.ModeCharDevice != 0 {
		r, err := iio.CreateReader(flags.JSONFile())
		if err != nil {
			fmt.Fprintln(os.Stderr, "create reader error:", err)
			os.Exit(1)
		}
		input = r
	}
	defer input.Close()

	table, err := json2csv.NewJSONFrom(input, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json error:", err)
		os.Exit(1)
	}

	out := os.Stdout
	if flags.CSVFile() != "" {
		w, err := iio.CreateWriter(flags.CSVFile())
		if err != nil {
			fmt.Fprintln(os.Stderr, "create writer error:", err)
			os.Exit(1)
		}
		out = w
	}
	defer out.Close()

	csv := json2csv.JSONtoCSV(table)
	err = json2csv.CSVTo(out, *csv)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json error:", err)
		os.Exit(1)
	}
}
