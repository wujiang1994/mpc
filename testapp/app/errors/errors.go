package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	status  int
	code    int
	message string
}

func (err Error) Error() string {
	return fmt.Sprintf(`{"status":%d,"code":%d,"message":%q}`, err.status, err.code, err.message)
}

func NewErrorWithStatus(status int, code int, message string) Error {
	return Error{status: status, code: code, message: message}
}

var (
	ErrForTest = NewErrorWithStatus(http.StatusInternalServerError, -1, "err for test")
)
