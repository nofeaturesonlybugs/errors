[![Documentation](https://godoc.org/github.com/nofeaturesonlybugs/errors?status.svg)](http://godoc.org/github.com/nofeaturesonlybugs/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/nofeaturesonlybugs/errors)](https://goreportcard.com/report/github.com/nofeaturesonlybugs/errors)
[![Build Status](https://travis-ci.com/nofeaturesonlybugs/errors.svg?branch=master)](https://travis-ci.com/nofeaturesonlybugs/errors)
[![codecov](https://codecov.io/gh/nofeaturesonlybugs/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/nofeaturesonlybugs/errors)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

golang package for enhanced errors that capture file and line number information, the call stack, relevant identifiers, and a handle back to the original error for type switching.

Replacement for fmt.Errorf(...):
```go
err := errors.Errorf("Failed to do something; here's a value %v",len(data))
fmt.Printf("%#v\n",err) // Prints full stack trace.
```

Wrap another function's error:
```go
_, err := net.Dial( "tcp", "localhost:8080" )
if err != nil {
    return errors.Go( err )
}
```

Obtain the original error in order to type switch:
```go
func connect() error {
    _, err := net.Dial( "tcp", "localhost:0" )
    if err != nil {
        return errors.Go( err )
    }
    return nil
}

err := connect()
switch errors.Original( err ).(type) {
    case *net.OpError:
        // We know the original type!
}
```

See the godoc link for more information.