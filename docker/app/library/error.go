package main

type LibraryError struct {
	statusCode int
	message    string
}

func NewLibraryError(statusCode int, message string) *LibraryError {
	return &LibraryError{
		statusCode: statusCode,
		message:    message,
	}
}
