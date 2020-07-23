package websearch

// The main web search error
type Error struct {
	error error
}

// Makes a new ProviderError
func NewError(err error) Error {
	return Error{
		error: err,
	}
}

func (e Error) Error() string {
	return e.error.Error()
}
