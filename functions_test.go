package errors_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/nofeaturesonlybugs/errors"
)

func ExampleOriginal() {
	// Demonstrates how to type switch on the original error when using this package.
	var err error

	// This test block ensures type switching works when nil is passed to Original.
	noError := func() error {
		return nil
	}
	err = noError()
	switch errors.Original(err).(type) {
	default:
		fmt.Printf("Works when error is nil.\n")
	}

	// This test block ensures type switching works when the error is Error from this package
	// but has a nil return value for Interface() method.
	errorf := func() error {
		return errors.Errorf("A formatted error!")
	}
	err = errorf()
	switch errors.Original(err).(type) {
	default:
		fmt.Printf("Works when error is error.Error but nil interface.\n")
	}

	// This test block shows the more practical usage of looking for a specifc error.
	netOpError := func() error {
		_, err := net.Dial("blurp", "localhost:0")
		if err != nil {
			return errors.Go(err)
		}
		return nil
	}
	err = netOpError()
	switch errors.Original(err).(type) {
	case *net.OpError:
		fmt.Printf("Got expected *net.OpErr\n")
	default:
		fmt.Printf("Expected *net.OpError; got %T\n", err)
	}

	// This final example shows an unwrapped error is returned as-is by Original.
	netOpErrorUnwrapped := func() error {
		_, err := net.Dial("blurp", "localhost:0")
		return err
	}
	err = netOpErrorUnwrapped()
	if originalErr := errors.Original(err); originalErr == err {
		fmt.Println("Original error returned when unwrapped.")
	}

	// Output: Works when error is nil.
	// Works when error is error.Error but nil interface.
	// Got expected *net.OpErr
	// Original error returned when unwrapped.
}

func TestIs(t *testing.T) {
	var ErrFoo = fmt.Errorf("this is a foo error")
	{
		e := ErrFoo
		if !errors.Is(e, ErrFoo) {
			t.Log("errors.Is(e, ErrFoo) returned false; true expected.")
			t.Fail()
		}
	}
	{
		e := errors.Go(ErrFoo)
		if !errors.Is(e, ErrFoo) {
			t.Log("errors.Is(e, ErrFoo) returned false; true expected.")
			t.Fail()
		}
	}
	{
		e := errors.Errorf("this is a foo error") // Same text but NOT ErrFoo
		if errors.Is(e, ErrFoo) {
			t.Log("errors.Is(e, ErrFoo) returned true; false expected.")
			t.Fail()
		}
	}
}
