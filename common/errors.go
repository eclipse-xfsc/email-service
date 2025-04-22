package common

import "fmt"

type EmailError struct {
	Msg  string
	Code int
}

func (e EmailError) Error() string {
	return fmt.Sprintf("error %v: %s", e.Code, e.Msg)
}
