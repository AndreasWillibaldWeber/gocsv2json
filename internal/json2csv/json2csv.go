package json2csv

import (
	"fmt"
	"io"

	"github.com/andreaswillibaldweber/gocsv2json/internal/iio"
	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
	v "github.com/andreaswillibaldweber/gocsv2json/internal/validater"
)

func NewJSONFrom(r io.Reader, header bool) (*m.JSON, error) {
	columns, err := iio.ReadJSON(r)
	if err != nil {
		return nil, fmt.Errorf("read JSON error: %w", err)
	}
	json := m.NewJSON(columns)
	if err := v.ValidateJSON(*json); err != nil {
		return nil, fmt.Errorf("validate JSON error: %w", err)
	}
	return json, nil
}

func JSONtoCSV(json *m.JSON) *m.CSV {
	return json.ToCSV()
}

func CSVTo(w io.Writer, csv m.CSV) error {
	return iio.WriteCSV(w, csv)
}
