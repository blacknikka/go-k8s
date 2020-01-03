package random

import "testing"

func TestGetRandomValue(t *testing.T) {
	rnd := NewRandom()
	got := rnd.GetRandomValue()

	if got <= 0 {
		t.Errorf("got %q", got)
	}
}
