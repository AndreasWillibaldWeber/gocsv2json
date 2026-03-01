package iio

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
)

func CreateWriter(p string) (*os.File, error) {
	out, err := os.Create(p)
	if err != nil {
		return nil, fmt.Errorf("create file error: %w", err)
	}
	return out, nil
}

func WriteJSON(w io.Writer, j m.JSON) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(j)
}

func WriteCSV(w io.Writer, csvIn m.CSV) error {
	writer := csv.NewWriter(w)
	if csvIn.Header() != nil {
		if err := writer.Write(rowToRecord(csvIn.Header())); err != nil {
			return err
		}
	}
	for _, row := range csvIn.Rows() {
		if err := writer.Write(rowToRecord(row)); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func rowToRecord(row m.Row) []string {
	record := make([]string, len(row))
	for i, cell := range row {
		record[i] = cell.ValueAsString()
	}
	return record
}
