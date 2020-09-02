package errors

import "fmt"

// Frame is a call stack frame.
type Frame struct {
	File     string
	Function string
	Line     int
}

// String returns the Frame as a string.
func (me Frame) String() string {
	return fmt.Sprintf("%v\n\t%v:%v", me.Function, me.Line, me.File)
}
