package merror

import (
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	err := errors.New("123")

	err = NewError(err)
	t.Log(err.Error())

	nerr := NewError(err)
	if nerr != err {
		t.Fail()
	}
}
