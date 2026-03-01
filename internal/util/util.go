package util

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
)

func WriteJSON(w io.Writer, j m.JSON) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(j)
}

func ReadCSV(r io.Reader) (m.Rows, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	cr.TrimLeadingSpace = true
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}
	return getRows(records), nil
}

func getRows(records [][]string) m.Rows {
	maxCols := maxColumns(records)
	rows := make(m.Rows, len(records))
	for i := range records {
		for j := range records[i] {
			val, kind := parseValueKind(records[i][j])
			rows[i] = append(rows[i], m.NewCell(val, kind))
		}
		for maxCols > len(rows[i]) {
			rows[i] = append(rows[i], m.NewCell("", m.KindString))
		}
	}
	return rows
}

func maxColumns(rs [][]string) int {
	maxCols := 0
	for i := range rs {
		cols := len(rs[i])
		if maxCols < cols {
			maxCols = cols
		}
	}
	return maxCols
}

func parseValueKind(s string) (any, m.Kind) {
	s = strings.TrimSpace(s)
	if s == "" {
		// return nil, m.KindNull
		return s, m.KindString
	}

	// time (RFC3339 / RFC3339Nano)
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, m.KindTime
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, m.KindTime
	}

	// duration (e.g. "150ms", "2h45m")
	if d, err := time.ParseDuration(s); err == nil {
		return d, m.KindDuration
	}

	// float (only if it looks like float, OR is NaN/Inf)
	looksFloat := strings.ContainsAny(s, ".eE") ||
		strings.EqualFold(s, "nan") ||
		strings.EqualFold(s, "+nan") ||
		strings.EqualFold(s, "-nan") ||
		strings.EqualFold(s, "inf") ||
		strings.EqualFold(s, "+inf") ||
		strings.EqualFold(s, "-inf") ||
		strings.EqualFold(s, "infinity") ||
		strings.EqualFold(s, "+infinity") ||
		strings.EqualFold(s, "-infinity")

	if looksFloat {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			// if you DON'T want NaN/Inf to count as float, remove this block
			if math.IsNaN(f) || math.IsInf(f, 0) {
				return f, m.KindFloat
			}
			return f, m.KindFloat
		}
	}

	// bool
	if b, err := strconv.ParseBool(s); err == nil {
		return b, m.KindBool
	}

	return s, m.KindString
}

func GetTypes(row m.Row) []string {
	var types []string
	for _, cell := range row {
		kind := cell.Kind()
		types = append(types, kind)
	}
	return types
}
