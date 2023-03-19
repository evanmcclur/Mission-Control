package jobs

import "fmt"

// Custom Http error that is returned when the status code is not 200
type HttpResponseError struct {
	StatusCode int
	Status     string
}

// Custom HTML parse error that is returned when an error has occurred while parsing/scraping HTML
type HTMLParseError struct {
	Job          string
	ErrorMessage string
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

// Constructs a new HTMLParseError
func NewHTMLParseError(job string, errorString string) error {
	return &HTMLParseError{
		Job:          job,
		ErrorMessage: errorString,
	}
}

func (e *HTMLParseError) Error() string {
	err := fmt.Sprintf("Job: %s Error: %s", e.Job, e.ErrorMessage)
	return err
}
