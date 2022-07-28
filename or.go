package or

import "testing"

var FormatText string = "accessing fail (%T): %+v"
var FormatTextCleanupE string = "accessing fail (%T), cleanup: %+v"

func Fatal[T any](ob T, err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()
		if err != nil {
			t.Fatalf(FormatText, ob, err)
		}
		return ob
	}
}

func FatalWithCleanup[T any](ob T, cleanup func(), err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()
		if cleanup != nil {
			t.Cleanup(cleanup)
		}
		if err != nil {
			t.Fatalf(FormatText, ob, err)
		}
		return ob
	}
}

func WithCleanup[T any](ob T, cleanup func()) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()
		if cleanup != nil {
			t.Cleanup(cleanup)
		}
		return ob
	}
}
func FatalWithCleanupE[T any](ob T, cleanup func() error, err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()
		if cleanup != nil {
			t.Cleanup(func() {
				t.Helper()
				if err := cleanup(); err != nil {
					t.Fatalf(FormatTextCleanupE, ob, err)
				}
			})
		}
		if err != nil {
			t.Fatalf(FormatText, ob, err)
		}
		return ob
	}
}

func WithCleanupE[T any](ob T, cleanup func() error) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()
		if cleanup != nil {
			t.Cleanup(func() {
				t.Helper()
				if err := cleanup(); err != nil {
					t.Fatalf(FormatTextCleanupE, ob, err)
				}
			})
		}
		return ob
	}
}
