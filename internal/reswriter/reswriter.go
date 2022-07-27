package reswriter

import (
	f "flights-assignment/internal/fail"
	"log"
	"net/http"
)

const (
	errLevel = "error"
)

type msg struct {
	Details string `json:"details"`
}

type marshaler interface {
	MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
}

type jsonResWriter struct {
	m marshaler
}

func New(m marshaler) *jsonResWriter {
	return &jsonResWriter{
		m: m,
	}
}

func (r *jsonResWriter) Write(w http.ResponseWriter, res interface{}, statusCode int) error {
	marshaledRes, err := r.m.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Printf("log level '%s' - reswriter Write: error marshaling response body '%s'",
			errLevel, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	/*
	** If "w.Write" returns an error, a good thing to do would be writing a
	** 4xx/5xx status code. The status code has already been written above.
	** A redundant "w.writeHeader" would be ignored.
	 */
	if _, err = w.Write(marshaledRes); err != nil {
		log.Printf("log level '%s' - reswriter Write: error writing response body '%s'",
			errLevel, err.Error())
	}
	return err
}

func (r *jsonResWriter) Err(w http.ResponseWriter, err error) {
	e, ok := err.(*f.Fail)
	if ok {
		// The nolint directives are specified because "r.Write" logs the error when it occurs
		switch e.Type {
		case f.ReqBody:
			r.Write(w, msg{Details: err.Error()}, http.StatusBadRequest) //nolint:errcheck
		case f.RespBody:
			r.Write(w, msg{Details: err.Error()}, http.StatusInternalServerError) //nolint:errcheck
		default:
			r.Write(w, msg{Details: err.Error()}, http.StatusInternalServerError) //nolint:errcheck
		}
	}
}
