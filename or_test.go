package or_test

import (
	"fmt"
	"testing"

	"github.com/podhmo/or"
)

type DB struct {
	Name string
}

func GetDB(name string) (*DB, error) {
	if name == "" {
		return nil, fmt.Errorf("must not be empty string")
	}
	return &DB{Name: name}, nil
}

func TestFatal(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := or.Fatal(GetDB("foo"))(t)
		want := &DB{Name: "foo"}
		if want, got := want.Name, got.Name; want != got {
			t.Errorf("Fatal() DB.Name want=%q but got=%q", want, got)
		}
	})
}
