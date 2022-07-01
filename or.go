package or

import "testing"

var FormatText string = "accessing fail (%T): %+v"

func Fatal[T any](ob T, err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		if err != nil {
			t.Fatalf(FormatText, ob, err)
		}
		return ob
	}
}