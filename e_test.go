package errors

import (
	"testing"
)

func TestNil(t *testing.T) {
	var e *e
	var p Error

	p = e
	if f := p.Stack(); len(f) != 0 {
		t.FailNow()
	} else if i := p.Interface(); i != nil {
		t.FailNow()
	} else if r := p.Tag("Hi", "There"); r != p {
		t.FailNow()
	} else if m := p.Error(); m != "" {
		t.FailNow()
	}

	p = Go(nil)
	if p != nil {
		t.FailNow()
	}
}
