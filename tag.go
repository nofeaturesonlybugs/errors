package errors

// tag represents a tag.
type tag struct {
	name, value string
}

func (me *tag) String() string {
	return me.name + "=" + me.value
}
