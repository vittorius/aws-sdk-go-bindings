package sns

const (

	// ErrEmptyParameter is used when a required parameter is empty
	ErrEmptyParameter = "EmptyParameter"

	// ErrNoPointerParameter is used when a parameter is expected to be a pointer but it wasn't
	ErrNoPointerParameter = "NoPointerParameter"

	// ErrPointerParameterNotAllowed is used when a parameter is expected to be not a pointer but it wasn't
	ErrPointerParameterNotAllowed = "PointerParameterNotAllowed"
)
