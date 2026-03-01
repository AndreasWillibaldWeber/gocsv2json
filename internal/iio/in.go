package iio

import (
	"encoding/csv"
	"encoding/json"
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

func ReadJSON(r io.Reader) (m.Columns, error) {
	var in m.JSON
	dec := json.NewDecoder(r)
	if err := dec.Decode(&in); err != nil {
		return nil, err
	}
	if in.Columns == nil {
		return make(m.Columns), nil
	}
	return in.Columns, nil
}
