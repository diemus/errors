package errors

import (
	"errors"
	"fmt"
	"testing"
)

func A() error {
	return Wrap(errors.New("test"))
}

func B() error {
	return Wrap(A())
}

func TestWrap(t *testing.T) {
	fmt.Println(B())
	fmt.Println(Cause(B()))
	if Cause(B()).Error() != "test" {
		t.Errorf("cause should equal origin error")
	}
}
