package marshaler

import "encoding/json"

type jsonMarshaler struct {
}

func New() *jsonMarshaler {
	return &jsonMarshaler{}
}

func (m *jsonMarshaler) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
