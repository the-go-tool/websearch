package errs

import "fmt"

// The not implemented error
type NotImplementedError struct {
	error error
}

// Makes a new NotImplementedError
func NewNotImplemented(err error) NotImplementedError {
	return NotImplementedError{
		error: err,
	}
}

func (e NotImplementedError) Error() string {
	return fmt.Sprintf("not implemented: %s", e.error.Error())
}
