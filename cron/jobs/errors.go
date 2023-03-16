package jobs

import "fmt"

// Custom Http error that is thrown when the status code is not 200
type HttpResponseError struct {
	StatusCode int
	Status     string
}

type HTMLParseError struct {
	Job   string
	Cause string
}

// Constructs a new HttpResponseError
func NewHttpResponseError(statusCode int, status string) error {
	return &HttpResponseError{
		StatusCode: statusCode,
		Status:     status,
	}
}

func (e *HttpResponseError) Error() string {
	err := fmt.Sprintf("Status code error: %d %s", e.StatusCode, e.Status)
	return err
}
