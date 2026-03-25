package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Add", func(t *testing.T) {
		want := 4
		got := Add(2, 2)
		assertCorrectMessage(t, got, want)
	})
}
