package reswriter

import (
	"net/http"
	f "flights-assignment/internal/fail"
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

func (r *jsonResWriter) Write(w http.ResponseWriter, res interface{}, statusCode int) {
	jsonResult, err := r.m.MarshalIndent(res, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	/*
	** If "w.Write" returns an error, a good thing to do would be writing
	** a 5xx status code. The status code has already been written above.
	** A redundant "w.writeHeader" would be ignored. This is why the error
	** is not checked and linting has been disabled.
	 */
	w.Write(jsonResult) //nolint:errcheck
}

func (r *jsonResWriter) Err(w http.ResponseWriter, err error) {
	e, ok := err.(*f.Fail)
	if ok {
		switch e.Type {
		case f.ReqBody:
			r.Write(w, msg{Details: err.Error()}, http.StatusBadRequest)
			return
		case f.RespBody:
			r.Write(w, msg{Details: err.Error()}, http.StatusInternalServerError)
			return
		default:
			r.Write(w, msg{Details: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
	r.Write(w, msg{Details: err.Error()}, http.StatusInternalServerError)
}
