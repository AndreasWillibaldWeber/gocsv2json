package models

import (
	"fmt"
)

type Cell struct {
	value any
	kind  Kind
}

func (c Cell) Value() any {
	return c.value
}

func (c Cell) ValueAsString() string {
	switch t := c.value.(type) {
	case nil:
		return ""
	case string:
		return t
	case []byte:
		return string(t)
	case fmt.Stringer:
		return t.String() // covers time.Time, time.Duration, etc.
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", t)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%g", t)
	case bool:
		return fmt.Sprintf("%t", t)
	default:
		return fmt.Sprintf("%v", t)
	}
}

func (c Cell) Kind() string {
	return c.kind.String()
}

func (c Cell) String() string {
	fmtStr := c.kind.FormatString() + "(%s)"
	return fmt.Sprintf(fmtStr, c.value, c.kind)
}

func NewCell(value any, kind Kind) Cell {
	return Cell{value, kind}
}
