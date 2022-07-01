package or_test

import (
	"fmt"
	"testing"

	"github.com/podhmo/or"
)

type DB struct {
	Name string
}

var errEmptyString = fmt.Errorf("must not be empty string")
var emptyString = ""

func GetDB(name string) (*DB, error) {
	if name == emptyString {
		return nil, errEmptyString
	}
	return &DB{Name: name}, nil
}

func GetDBWithCleanup(name string) (*DB, func(), error) {
	if name == emptyString {
		return nil, nil, errEmptyString
	}
	return &DB{Name: name}, func() {}, nil
}

type FakeTB struct {
	testing.TB

	fmt    string
	args   []interface{}
	called bool
}

func (tb *FakeTB) Fatalf(fmt string, args ...interface{}) {
	tb.fmt = fmt
	tb.args = args
	tb.called = true
}

func TestFatal(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := or.Fatal(GetDB("foo"))(t) // ok

		want := &DB{Name: "foo"}
		if want, got := want.Name, got.Name; want != got {
			t.Errorf("Fatal(), DB.Name want=%q but got=%q", want, got)
		}
	})

	t.Run("ng", func(t *testing.T) {
		fake := &FakeTB{TB: t}
		_ = or.Fatal(GetDB(emptyString))(fake) // ng

		if !fake.called {
			t.Fatalf("Fatal(), must be failed, but successed")
		}

		want := fmt.Sprintf(or.FormatText, &DB{}, errEmptyString) // %T %+v
		got := fmt.Sprintf(fake.fmt, fake.args...)
		if want != got {
			t.Errorf("Fatal(), error message is want=%q but got=%q", want, got)
		}
	})
}

func TestFatalWithCleanup(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got, cleanup := or.FatalWithCleanup(GetDBWithCleanup("foo"))(t) // ok
		defer cleanup()

		want := &DB{Name: "foo"}
		if want, got := want.Name, got.Name; want != got {
			t.Errorf("Fatal(), DB.Name want=%q but got=%q", want, got)
		}
	})
}
