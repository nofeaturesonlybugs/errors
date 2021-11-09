package errors

import (
	"errors"
	"runtime"
)

// Is reports whether any error in err's chain matches target.
//
// If err is an Error returned from this package then Is() tests the error
// returned from Original(err).
//
// Internally it calls down to errors.Is() from the standard errors package.
func Is(err error, target error) bool {
	switch e := err.(type) {
	case *e:
		return errors.Is(e.original, target)
	default:
		return errors.Is(err, target)
	}
}

// Original returns a non-nil interface if the incoming type is an Error from this package.
func Original(err error) interface{} {
	if e, ok := err.(Error); ok {
		return e.Interface()
	} else if err != nil {
		return err
	}
	return nil
}

// Stack returns the call stack frames of the caller.
func Stack() []Frame {
	var frames []Frame
	var ptrs []uintptr
	skip := 1
	for pc, _, _, ok := runtime.Caller(skip); ok; pc, _, _, ok = runtime.Caller(skip) {
		skip++
		ptrs = append(ptrs, pc)
	}
	//
	stack := runtime.CallersFrames(ptrs)
	for frame, ok := stack.Next(); ok; frame, ok = stack.Next() {
		frames = append(frames, Frame{File: frame.File, Function: frame.Function, Line: frame.Line})
	}
	//
	return frames
}
