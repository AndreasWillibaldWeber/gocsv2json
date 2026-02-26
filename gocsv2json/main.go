package main

import (
	"fmt"
	"os"

	"github.com/andreaswillibaldweber/csv2json/internal/csv2json"
	"github.com/andreaswillibaldweber/csv2json/internal/csv2json/cli"
)

func main() {
	flags := cli.ParseFlags()
	fmt.Printf("Flags>> %s \n\n", flags)

	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin stat error:", err)
		os.Exit(2)
	}
	if fi.Mode()&os.ModeCharDevice != 0 {
		fmt.Fprintln(os.Stderr, "No stdin pipe detected. Usage: cat file.csv | csv2json")
		os.Exit(2)
	}

	table, err := csv2json.NewCSVFrom(os.Stdin, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "csv error:", err)
		os.Exit(1)
	}

	// table is your parsed "table" (rows x columns)
	fmt.Printf("Rows: %d\n", table.Len())
	fmt.Println(table)

	out, err := csv2json.CreateWriter("examples/time_series.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "create writer error:", err)
		os.Exit(1)
	}
	defer out.Close()

	json := csv2json.CSVtoJSON(table)
	err = csv2json.JSONTo(out, *json)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json error:", err)
		os.Exit(1)
	}

	out = os.Stdout
	err = csv2json.JSONTo(out, *json)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json error:", err)
		os.Exit(1)
	}
}
