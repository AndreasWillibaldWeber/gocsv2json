package models

type JSON struct {
	Columns Columns `json:"columns"`
}

func (j *JSON) ToCSV() *CSV {
	return NewEmptyCSV()
}

func NewJSON(columns Columns) *JSON {
	return &JSON{columns}
}

func NewEmptyJSON() *JSON {
	return &JSON{
		Columns: make(Columns),
	}
}
