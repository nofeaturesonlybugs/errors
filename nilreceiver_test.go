package errors_test

import (
	"fmt"
	"github.com/nofeaturesonlybugs/errors"
)

type A struct{}

func (a *A) String() (string, error) {
	if a == nil {
		return "", errors.NilReceiver().Type(a)
	}
	return "Hello, World!", nil
}

func ExampleError_nilReceiver() {
	var a *A

	a = &A{}
	if s, err := a.String(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	a = nil
	if s, err := a.String(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	// Output: Hello, World!
	// Nil receiver; type=*errors_test.A
}
