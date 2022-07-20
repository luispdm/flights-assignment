package handler

import (
	"io"
	"log"
	"net/http"
	"flights-assignment/internal/differ"
	"flights-assignment/internal/fail"
	"flights-assignment/internal/tracker"
)

const (
	errLevel   = "error"
	traceLevel = "trace"
)

type decoder interface {
	Decode(r io.Reader, v interface{}) error
}

type resWriter interface {
	Write(w http.ResponseWriter, res interface{}, statusCode int)
	Err(w http.ResponseWriter, err error)
}

func PostTrack(r resWriter, dec decoder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		var flights []tracker.Flight
		err := dec.Decode(req.Body, &flights)
		if err != nil {
			wrappedErr := fail.New(fail.ReqBody, err.Error())
			log.Printf("log level '%s' - PostTrack: %v", errLevel, wrappedErr)
			r.Err(w, wrappedErr)
			return
		}
		log.Printf("log level '%s' - PostTrack: request body decoded successfully", traceLevel)

		tracked := tracker.New(flights, differ.New()).Track()
		r.Write(w, tracked, http.StatusOK)
		if err != nil {
			wrappedErr := fail.New(fail.RespBody, err.Error())
			log.Printf("log level '%s' - PostTrack: %v", errLevel, wrappedErr)
			r.Err(w, wrappedErr)
			return
		}
		log.Printf("log level '%s' - PostTrack: response body '%v' written with status code '%d'",
			traceLevel, tracked, http.StatusOK)
	}
}
