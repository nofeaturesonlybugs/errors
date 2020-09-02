// Package errors provides enhanced error reporting and specific error types.
//
// This package provides errors with more information about when and where they occurred; the
// call to Printf will print file name, line number, and stack information.
//	err := errors.Errorf("An error occurred because len() == %v", len(something))
//	fmt.Printf("%#v\n",err) // Prints full stack trace.
//
// Go() can wrap an existing error within the Error interface defined in this package.
//	err := someApi()
//	if err != nil {
//		return errors.Go(err)
//	}
//
// Utility Methods
//
// For receivers that are pointers:
//	func (me *Type) Do() error {
//		if me == nil {
//			return errors.NilReceiver().Type(me)
//		}
//	}
//
// For required arguments that are nil:
//	func Make( arg *Conf ) (*Type, error) {
//		if arg == nil {
//			return nil, errors.NilArgument( "arg" ).Type( arg );
//		}
//	}
//
// Original Error
//
// It may be helpful for your error handling code to type switch on the original error when
// wrapping it within errors.Go(); see documentation and example for Original().
package errors
