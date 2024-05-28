package ghost

type HttpError interface {
	Message() string
}

type NotFoundError struct {
	message string
}

func (e NotFoundError) Message() string {
	return e.message
}
