# or

shorthand for testing in go

## required

go 1.18 (for generics)

## how to use

`Fatal()` function is a shorthand for the factory function returning with error value (e.g. `func(...) (T, error)`)

Use it like this.

```go
import (
	"fmt"
	"testing"

	"github.com/podhmo/or"
)

type Foo struct{}

func getFoo() (*Foo, error) { return nil, fmt.Errorf("hmm") }

type Bar struct{ Foo *Foo }

func getBar(foo *Foo) (*Bar, error) { return nil, fmt.Errorf("hmm") }

func TestIt(t *testing.T) {
	foo := or.Fatal(getFoo())(t)
	bar := or.Fatal(getBar(foo))(t)

	_ = bar // doSomething
}
```

The result is here.

```console
$ go test
--- FAIL: TestIt (0.00s)
    or.go:10: accessing fail (*or_test.Foo): hmm
``` 

### why not `or.Fatal(t, getFoo()` ?

Because two reasons.

- for go's multiple values specification
- for go's generics specification