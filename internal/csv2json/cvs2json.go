package csv2json

import (
	"fmt"
	"io"

	"github.com/andreaswillibaldweber/gocsv2json/internal/iio"
	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
	v "github.com/andreaswillibaldweber/gocsv2json/internal/validater"
)

func NewCSVFrom(r io.Reader, header bool) (*m.CSV, error) {
	rows, err := iio.ReadCSV(r)
	if err != nil {
		return nil, fmt.Errorf("read CSV error: %w", err)
	}
	csv := m.NewCSV(rows, header)
	if err := v.ValidateCSV(*csv); err != nil {
		return nil, fmt.Errorf("validate CSV error: %w", err)
	}
	return csv, nil
}

func CSVtoJSON(csv *m.CSV) *m.JSON {
	return csv.ToJSON()
}

func JSONTo(w io.Writer, json m.JSON) error {
	return iio.WriteJSON(w, json)
}
