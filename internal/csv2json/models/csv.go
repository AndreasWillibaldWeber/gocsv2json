package models

import (
	"fmt"
	"strings"
)

type CSV struct {
	header Row
	rows   Rows
}

func (c CSV) NonEmptyHeader() Row {
	header := Row{}
	for col, h := range c.header {
		val := strings.TrimSpace(h.ValueAsString())
		val = strings.ReplaceAll(val, " ", "_")
		if h.kind.String() == "STRING" && val == "" {
			colHeader := fmt.Sprintf("col%d", col)
			header = append(header, NewCell(colHeader, h.kind))
			continue
		}
		header = append(header, NewCell(val, h.kind))
	}
	return header
}

func (c CSV) Header() Row {
	return c.header
}

func (c CSV) Rows() Rows {
	return c.rows
}

func (c CSV) Len() int {
	rows := len(c.rows)
	if c.header != nil {
		rows += 1
	}
	return rows
}

func (c CSV) ToJSON() *JSON {
	json := NewEmptyJSON()
	for _, h := range c.NonEmptyHeader() {
		json.Columns[h.ValueAsString()] = make(Column, len(c.rows))
	}
	for i, r := range c.rows {
		for j, h := range c.NonEmptyHeader() {
			json.Columns[h.ValueAsString()][i] = r[j].Value()
		}
	}
	return json
}

func (c CSV) String() string {
	if c.header != nil {
		return fmt.Sprintf("%s\n%s", c.header, c.rows)
	}
	return fmt.Sprintf("%s", c.rows)
}

func NewCSV(rows Rows, header bool) *CSV {

	if len(rows) == 0 {
		return &CSV{
			header: nil,
			rows:   nil,
		}
	}

	if header {
		if len(rows) > 1 {
			return &CSV{
				header: rows[0],
				rows:   rows[1:],
			}
		}
		return &CSV{
			header: rows[0],
			rows:   nil,
		}
	}

	return &CSV{
		header: nil,
		rows:   rows,
	}
}
