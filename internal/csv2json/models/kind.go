package models

const (
	KindNull Kind = iota
	KindTime
	KindDuration
	KindBool
	KindInt
	KindFloat
	KindString
)

type Kind int

func (k Kind) FormatString() string {
	switch k {
	case KindBool:
		return "%t"
	case KindInt:
		return "%d"
	case KindFloat:
		return "%g"
	case KindString:
		return "%s"
	case KindDuration:
		return "%s"
	default:
		return "%v"
	}
}

func (k Kind) String() string {
	switch k {
	case KindNull:
		return "NULL"
	case KindTime:
		return "TIME"
	case KindDuration:
		return "DURATION"
	case KindBool:
		return "BOOL"
	case KindInt:
		return "INT"
	case KindFloat:
		return "FLOAT"
	case KindString:
		return "STRING"
	default:
		return "UNKNOWN"
	}
}
