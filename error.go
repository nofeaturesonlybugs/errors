package errors

import "fmt"

// Error is the interface returned by functions in this package.
type Error interface {
	// Error returns a string value of the Error and satisfies built-in error interface.
	Error() string
	// Format implements fmt.Formatter interface; %#v, %+v will print error information plus call stack.
	Format(fmt.State, rune)
	// Interface returns the underlying error in order to support type switching; this allows error types from
	// other packages to be wrapped with Go() but still type switch on the underlying type.
	Interface() interface{}
	// Stack returns slice of Frame objects to represent the stack frame.
	Stack() []Frame
	// Tag adds extra information to the error and returns the type for method chaining.
	Tag(name, value string) Error
	// Type adds a special Tag() with type=%T information.
	Type(interface{}) Error
}

// AlreadyStarted returns an error representing a service that has already started and can not start again.
func AlreadyStarted() Error {
	return make("Service already started", nil)
}

// Errorf is similar to fmt.Errorf except it returns the enhanced Error interface from this package.
func Errorf(format string, args ...interface{}) Error {
	return make(fmt.Sprintf(format, args...), nil)
}

// Go returns an Error if the passed `err` is non-nil; the purpose of this function is to wrap contextual
// information, such as call stack, to the initial error.
func Go(err error) Error {
	if err != nil {
		// If err implements Error then we can return it without making a new one so as not
		// to lose the original call stack, original error, etc.
		switch typed := err.(type) {
		case Error:
			return typed
		default:
			return make(err.Error(), err)
		}
	}
	return nil
}

// NilArgument creates an error when a function argument is nil.
// func MyFunc(a *int) (int, error) {
// 	if a == nil {
// 		return 0, NilArgument("a").Type(a)
// 	}
// 	return *a, nil
// }
func NilArgument(ident string) Error {
	return make("Nil argument", nil).Tag("identifier", ident)
}

// NilMember creates an error when a struct member is nil.
// func (me SomeType) DoIt() error {
// 	if me.member == nil {
// 		return NilMember("me.member").Type(me.member)
// 	}
// 	return nil
// }
func NilMember(ident string) Error {
	return make("Nil member", nil).Tag("member", ident)
}

// NilReceiver creates an error when a receiver is nil.
// func (me *SomeType) DoIt() error {
// 	if me == nil {
// 		return NilReceiver()
// 	}
// 	return nil
// }
func NilReceiver() Error {
	return make("Nil receiver", nil)
}
