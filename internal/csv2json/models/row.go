package models

import "fmt"

type Row []Cell

func (r Row) String() string {
	out := ""
	for i, c := range r {
		if i == 0 {
			out += fmt.Sprintf("%s", c)
			continue
		}
		out += fmt.Sprintf(",%s", c)
	}
	return out
}

type Rows []Row

func (rs Rows) String() string {
	out := ""
	for i, r := range rs {
		if i != len(rs)-1 {
			out += fmt.Sprintf("%s\n", r)
			continue
		}
		out += fmt.Sprintf("%s", r)
	}
	return out
}
