package or_test

import (
	"fmt"
	"testing"

	"github.com/podhmo/or"
)

type Foo struct{}

func getFoo() (*Foo, error) { return &Foo{}, nil }

type Bar struct{ Foo *Foo }

func getBar(foo *Foo) (*Bar, error) { return nil, fmt.Errorf("hmm") }

type Boo struct{ Foo *Foo }

func getBoo(foo *Foo) (*Boo, func(), error) { return nil, nil, fmt.Errorf("hmm") }

func TestFatal(t *testing.T) {
	foo := or.Fatal(getFoo())(t)
	bar := or.Fatal(getBar(foo))(t)

	_ = bar // doSomething
}

func TestFatalWithCleanup(t *testing.T) {
	foo := or.Fatal(getFoo())(t)
	boo, cleanup := or.FatalWithCleanup(getBoo(foo))(t)
	defer cleanup()

	_ = boo // doSomething
}
