package errs

import "fmt"

// The bad request error
type BadRequestError struct {
	error error
}

// Makes a new BadRequestError
func NewBadRequestError(err error) BadRequestError {
	return BadRequestError{
		error: err,
	}
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("bad request: %s", e.error.Error())
}
