package errs

import "fmt"

// The ip banned error
type IPBannedError struct {
	error error
}

// Makes a new IPBannedError
func NewIPBanned(err error) IPBannedError {
	return IPBannedError{
		error: err,
	}
}

func (e IPBannedError) Error() string {
	return fmt.Sprintf("ip banned: %s", e.error.Error())
}
