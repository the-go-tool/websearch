package errs

import "fmt"

// The technical error
type TechnicalError struct {
	error error
}

// Makes a new NewTechnical
func NewTechnical(err error) TechnicalError {
	return TechnicalError{
		error: err,
	}
}

func (e TechnicalError) Error() string {
	return fmt.Sprintf("technical: %s", e.error.Error())
}
