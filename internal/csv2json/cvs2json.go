package csv2json

import (
	"fmt"
	"io"
	"os"

	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
	u "github.com/andreaswillibaldweber/gocsv2json/internal/util"
	v "github.com/andreaswillibaldweber/gocsv2json/internal/validater"
)

func NewCSVFrom(r io.Reader, header bool) (*m.CSV, error) {
	rows, err := u.ReadCSV(r)
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
	return u.WriteJSON(w, json)
}

func CreateWriter(p string) (*os.File, error) {
	out, err := os.Create(p)
	if err != nil {
		return nil, fmt.Errorf("create file error: %w", err)
	}
	return out, nil
}

func CreateReader(p string) (*os.File, error) {
	in, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}
	return in, nil
}
