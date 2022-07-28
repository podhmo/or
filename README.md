# or

shorthand for testing in go

## required

go 1.18 (for generics)

## how to use

`Fatal()` function is a shorthand for the factory function returning with error value (e.g. `func(...) (T, error)`)

Use it like this.

```go
package main

import (
	"fmt"
	"testing"

	"github.com/podhmo/or"
)

func TestFatal(t *testing.T) {
	foo := or.Fatal(getFoo())(t)
	bar := or.Fatal(getBar(foo))(t)

	_ = bar // doSomething
}

func TestFatalWithCleanup(t *testing.T) {
	foo := or.Fatal(getFoo())(t)
	boo := or.FatalWithCleanup(getBoo(foo))(t)

	_ = boo // doSomething
}

type Foo struct{}

func getFoo() (*Foo, error) { return &Foo{}, nil }

type Bar struct{ Foo *Foo }

func getBar(foo *Foo) (*Bar, error) { return nil, fmt.Errorf("hmm") }

type Boo struct{ Foo *Foo }

func getBoo(foo *Foo) (*Boo, func(), error) { return nil, nil, fmt.Errorf("hmm") }
```

The result is here.

```console
$ go test
--- FAIL: TestFatal (0.00s)
    example_test.go:24: accessing fail (*main.Bar): hmm
--- FAIL: TestFatalWithCleanup (0.00s)
    example_test.go:31: accessing fail (*main.Boo): hmm
FAIL
exit status 1
FAIL	github.com/podhmo/or/testdata	0.003s
``` 

### why not `or.Fatal(t, getFoo()` ?

There are two reasons.

- go's multiple values specification
- go's generics specification