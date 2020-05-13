# errors
the most easiest way to print stacktrace for error in golang

## Required Go Version

    golang >= 1.13

## Usage

import 

```go
import "github.com/diemus/errors"
```

Wrap an error when you return it, then you will get an stacktrace when you print it

```go
func test() error {
	err := errors.New("test")
	return errors.Wrap(err)
}

func test2() error {
	_ ,err := someLib.do()
	return errors.Wrap(err)
}


func test1(err error) {
    log.Println(err)
}
```

use Cause() to get the original error when you don't want show the full detail of stacktrace

```go
func test2(err error) {
    log.Println(errors.Cause(err))
}
```

the other function in standard library "errors" are wrap by this library, so you can replace the standard library with this package

```go
import "github.com/diemus/errors"

func test3(err error) {
	errors.New("test")
	errors.Is(err,err)
	errors.As(err,err)
}
```