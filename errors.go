package godesu

import "fmt"

type ErrNotOK struct {
	code int
}

type ErrThreadNotFound struct {
	number int
}

func (e ErrNotOK) Error() string {
	return fmt.Sprintf("http status not ok: %d", e.code)
}

func (e ErrThreadNotFound) Error() string {
	return fmt.Sprintf("thread %d not found", e.number)
}