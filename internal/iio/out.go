package iio

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	m "github.com/andreaswillibaldweber/gocsv2json/internal/models"
)

func CreateWriter(p string) (*os.File, error) {
	out, err := os.Create(p)
	if err != nil {
		return nil, fmt.Errorf("create file error: %w", err)
	}
	return out, nil
}

func WriteJSON(w io.Writer, j m.JSON) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(j)
}
