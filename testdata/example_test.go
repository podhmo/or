package or_test

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
