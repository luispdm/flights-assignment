package decoder

import (
	"encoding/json"
	"io"
	"log"
)

type jsonDecoder struct {
}

func New() *jsonDecoder {
	return &jsonDecoder{}
}

func (d *jsonDecoder) Decode(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		log.Printf("log level 'error' - decoder Decode: error decoding object '%s'",
			err.Error())
	}
	return err
}
