package iio

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
	u "github.com/andreaswillibaldweber/gocsv2json/internal/util"
)

func CreateReader(p string) (*os.File, error) {
	in, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}
	return in, nil
}

func ReadCSV(r io.Reader) (m.Rows, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	cr.TrimLeadingSpace = true
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}
	return u.GetRows(records), nil
}
