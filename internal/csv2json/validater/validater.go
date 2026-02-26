package validater

import (
	"fmt"

	m "github.com/andreaswillibaldweber/gocsv2json/internal/csv2json/models"
	u "github.com/andreaswillibaldweber/gocsv2json/internal/csv2json/util"
)

func ValidateCSV(csv m.CSV) error {
	if err := validateHeader(csv.Header()); err != nil {
		return err
	}
	if err := validateRows(csv.Rows()); err != nil {
		return err
	}
	return nil
}

func validateHeader(header m.Row) error {
	return missingHeader(header)
}

func validateRows(rows m.Rows) error {
	if err := sameLength(rows); err != nil {
		return err
	}
	if err := consistentKinds(rows); err != nil {
		return err
	}
	return nil
}

func missingHeader(header m.Row) error {
	if header == nil {
		return fmt.Errorf("missing header")
	}
	return nil
}

func consistentKinds(rows m.Rows) error {
	if len(rows) > 0 {
		types := u.GetTypes(rows[0])
		for _, row := range rows {
			for i, cell := range row {
				if cell.Kind() != types[i] {
					return fmt.Errorf("inconsistent types in column %d: expected %s, got %s", i, types[i], cell.Kind())
				}
			}
		}
	}
	return nil
}

func sameLength(rows m.Rows) error {
	cols := -1
	for i, row := range rows {
		if cols == -1 {
			cols = len(row)
		}
		if len(row) != cols {
			return fmt.Errorf("row %d has %d columns, expected %d", i, len(row), cols)
		}
	}
	return nil
}

func ValidateJSON(json m.JSON) error {
	// Implement your validation logic here
	return nil
}
