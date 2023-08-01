package set_test

import (
	"testing"

	"github.com/s0rg/set"
)

func TestSet(t *testing.T) {
	t.Parallel()

	s := make(set.Set[string])

	const (
		val1 = "a"
		val2 = "b"
		val3 = "c"
	)

	s.Add(val1)
	s.Add(val2)
	s.Load(val1, val2)

	if len(s) != 2 {
		t.Error("unexpected length")
	}

	if len(s.List()) != 2 {
		t.Error("unexpected list length")
	}

	if !s.Has(val1) {
		t.Error("no val1")
	}

	if s.Has(val3) {
		t.Error("has val3")
	}

	if !s.TryAdd(val3) {
		t.Error("TryAdd(val3) == false")
	}

	if s.TryAdd(val3) {
		t.Error("TryAdd(val3) == true")
	}
}
