package decoder

import (
	"encoding/json"
	"io"
)

type jsonDecoder struct {
}

func New() *jsonDecoder {
	return &jsonDecoder{}
}

func (d *jsonDecoder) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
