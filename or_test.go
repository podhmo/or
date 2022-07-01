package or_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
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
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Fatal() mismatch (-want +got):\n%s", diff)
		}
	})
}
