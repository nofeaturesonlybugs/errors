package errors_test

import (
	"fmt"
	"github.com/nofeaturesonlybugs/errors"
)

type B struct {
	V *int
}

func (b *B) DoIt() error {
	if b == nil {
		return errors.NilReceiver()
	} else if b.V == nil {
		return errors.NilMember("V").Type(b.V)
	}
	fmt.Println("Did it!")
	return nil
}

func (b *B) String() (string, error) {
	if b == nil {
		return "", errors.NilReceiver().Type(b)
	}
	return "Hello, World!", nil
}

func ExampleError_nilMember() {
	var b *B

	b = &B{}
	if s, err := b.String(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	if err := b.DoIt(); err != nil {
		fmt.Println(err)
	}

	b = nil
	if s, err := b.String(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	// Output: Hello, World!
	// Nil member; member=V; type=*int
	// Nil receiver; type=*errors_test.B
}
