package pcerror

import "errors"

type Error error

var (
	ErrorValueNotFound Error = errors.New("value not found")
)
