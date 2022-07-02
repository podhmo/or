package or

import "testing"

var FormatText string = "accessing fail (%T): %+v"

func Fatal[T any](ob T, err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()
		if err != nil {
			t.Fatalf(FormatText, ob, err)
		}
		return ob
	}
}

func FatalWithCleanup[T any](ob T, cleanup func(), err error) func(t testing.TB) (T, func()) {
	return func(t testing.TB) (T, func()) {
		t.Helper()
		if err != nil {
			t.Fatalf(FormatText, ob, err)
		}
		if cleanup == nil {
			cleanup = func() {}
		}
		return ob, cleanup
	}
}
