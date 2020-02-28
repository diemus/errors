package errors

import (
	"errors"
	"fmt"
	"runtime"
)

//wrap an error to record stacktrace
func Wrap(err error) error {
	if pc, file, line, ok := runtime.Caller(1); ok {
		f := runtime.FuncForPC(pc)
		var funcName string
		if f != nil {
			funcName = f.Name()
		}
		return fmt.Errorf("%w\n%s\n\t%s:%d", err, funcName, file, line)
	}
	return err
}

//return the original error without stacktrace
func Cause(err error) error {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}
	return err
}

func New(text string) error {
	return errors.New(text)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
