package fail

import (
	"fmt"
)

const (
	ReqBody        = "ReqBody"
	reqBodyMessage = "error reading request body"
	RespBody       = "RespBody"
	resBodyMessage = "error writing response body"

	unknownErrorMessage = "something went wrong"
)

type Fail struct {
	e    error
	Type string
}

func New(failType string, details string) *Fail {
	var err error

	switch failType {
	case ReqBody:
		err = fmt.Errorf("%s: '%s'", reqBodyMessage, details)
	case RespBody:
		err = fmt.Errorf("%s: '%s'", resBodyMessage, details)
	default:
		err = fmt.Errorf("%s: '%s'", unknownErrorMessage, details)
	}
	return &Fail{
		e:    err,
		Type: failType,
	}
}

func (f *Fail) Error() string {
	return f.e.Error()
}
