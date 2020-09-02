package errors

import (
	"fmt"
	"io"
	"strings"
)

// e is the internal type that fulfills the Error interface.
type e struct {
	message  string
	original interface{}
	stack    []Frame
	//
	tags []*tag
}

// make creates a new *e
func make(message string, original error) *e {
	rv := &e{
		message:  message,
		original: original,
		tags:     []*tag{}}
	// Get call stack; first two levels are removed to prune the calls into this package.
	rv.stack = Stack()[2:]
	// If original is nil then we are are the original type
	if original == nil {
		rv.original = rv
	}
	return rv
}

func (me *e) Error() string {
	if me != nil {
		items := []string{me.message}
		for _, tag := range me.tags {
			items = append(items, tag.String())
		}
		return strings.Join(items, "; ")
	}
	return ""
}

func (me *e) Format(state fmt.State, verb rune) {
	if me != nil {
		switch verb {
		case 'v':
			fallthrough
		case 's':
			io.WriteString(state, me.Error())
			if state.Flag('+') || state.Flag('#') {
				lines := []string{}
				for _, frame := range me.Stack() {
					lines = append(lines, frame.String())
				}
				if len(lines) > 0 {
					io.WriteString(state, "\n"+strings.Join(lines, "\n"))
				}
			}
		}
	}
}

func (me *e) Interface() interface{} {
	if me != nil {
		return me.original
	}
	return nil
}

func (me *e) Stack() []Frame {
	if me == nil {
		return []Frame{}
	}
	return me.stack
}

func (me *e) Tag(name, value string) Error {
	if me != nil {
		me.tags = append(me.tags, &tag{name: name, value: value})
	}
	return me
}

func (me *e) Type(v interface{}) Error {
	return me.Tag("type", fmt.Sprintf("%T", v))
}
