package errors_test

import (
	"fmt"
	"github.com/nofeaturesonlybugs/errors"
	"net"
	"testing"
)

func ExampleError_nilArgument() {
	type T struct{}

	fn := func(arg *T) error {
		if arg == nil {
			return errors.NilArgument("arg").Type(arg)
		}
		return nil
	}

	if err := fn(nil); err != nil {
		fmt.Println(err)
		_ = fmt.Sprintf("%#v", err)
	}
	// Output: Nil argument; identifier=arg; type=*errors_test.T
}

func TestAlreadyStarted(t *testing.T) {
	err := errors.AlreadyStarted()
	if err.Error() != "Service already started" {
		t.FailNow()
	}
}

func TestGo(t *testing.T) {
	{
		err := errors.AlreadyStarted()
		wrapped := errors.Go(err)
		if err != wrapped {
			t.FailNow()
		}
	}
	{
		_, netop := net.Dial("blurp", "localhost:0")
		wrapped := errors.Go(netop)
		if netop == wrapped {
			t.FailNow()
		}
		if errors.Original(wrapped) != netop {
			t.FailNow()
		}
	}
}
