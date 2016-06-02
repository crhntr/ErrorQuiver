package ErrorQuiver

import (
	"bytes"
	"errors"
)

type ErrorQuiver []error

const ErrorDelimiter = "; "

func New() *ErrorQuiver {
	return new(ErrorQuiver)
}

// Allows adding nil.
func (es *ErrorQuiver) Add(err error) {
	if err != nil {
		*es = append(*es, err)
	}
}

// use instead of
//
//     errs.Add(errors.New("something went wrong"))
//
func (es *ErrorQuiver) AddNew(errStr string) {
	if errStr != "" {
		*es = append(*es, errors.New(errStr))
	}
}

func (es *ErrorQuiver) Len() int {
	return len((*es))
}

// Does not preserve order
// Should not have memory leak
// HAS NO TESTS
//func (es *ErrorSlice) Get() error {
//	err := (*es)[len((*es))-1]
//	(*es)[len((*es))-1] = nil
//	(*es) = (*es)[:len((*es))-1]
//	return err
//}

// Will cause recursion when ErrorSlice contains ErrorSlices
func (es *ErrorQuiver) Error() string {
	buffer := bytes.Buffer{}
	for _, err := range *es {
		buffer.WriteString(err.Error() + ErrorDelimiter)
	}
	return buffer.String()
}

// This function can be used to check if the function should return.
// Instead of
//
//     if err != nil {
//         return err
//     }
//
// use
//
//     if errs.HaveBeenAdded() {
//         return errs
//     }
//
func (es *ErrorQuiver) HaveBeenAdded() bool {
	if len((*es)) > 0 {
		return true
	}
	return false
}

// Has no tests
func (es *ErrorQuiver) Return() error {
	if len((*es)) > 0 {
		return es
	}
	return nil
}
