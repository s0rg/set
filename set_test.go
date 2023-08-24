package set_test

import (
	"sort"
	"testing"

	"github.com/s0rg/set"
)

func testSet(t *testing.T, s set.Set[string]) {
	t.Helper()

	const (
		val1 = "a"
		val2 = "b"
		val3 = "c"
	)

	s.Add(val1)
	s.Add(val2)

	set.Load(s, val1, val2)

	if s.Len() != 2 {
		t.Error("unexpected length")
	}

	if len(set.ToSlice(s)) != 2 {
		t.Error("unexpected slice length")
	}

	if !s.Has(val1) {
		t.Error("no val1")
	}

	if s.Has(val3) {
		t.Error("has val3")
	}

	if !s.Add(val3) {
		t.Error("TryAdd(val3) == false")
	}

	if s.Add(val3) {
		t.Error("TryAdd(val3) == true")
	}

	var count int

	s.Iter(func(_ string) bool {
		count++

		return true
	})

	if count != s.Len() {
		t.Error("iter - count mismatch")
	}

	c := s.Clone()

	s.Del(val2)

	if s.Has(val2) {
		t.Error("has val2")
	}

	s.Clear()

	if s.Len() > 0 {
		t.Error("non-empty")
	}

	if !c.Has(val3) {
		t.Error("cloned missing val3")
	}

	c.Del(val2)
	c.Del(val3)
	c.Del(val3) // double-deletion test

	var res string

	c.Iter(func(v string) bool {
		res = v

		return false
	})

	if res != val1 {
		t.Fail()
	}
}

func testPop(t *testing.T, s set.Set[int]) {
	t.Helper()

	if _, ok := s.Pop(); ok {
		t.Fail()
	}

	s.Add(1)

	if v, ok := s.Pop(); !ok || v != 1 {
		t.Fail()
	}

	if s.Len() != 0 {
		t.Fail()
	}
}

func TestUnordered(t *testing.T) {
	t.Parallel()

	testSet(t, set.NewUnordered[string]())
	testPop(t, set.NewUnordered[int]())
}

func TestOrdered(t *testing.T) {
	t.Parallel()

	testSet(t, set.NewOrdered[string]())
	testPop(t, set.NewOrdered[int]())
}

func TestOrderedOrder(t *testing.T) {
	t.Parallel()

	s := set.NewOrdered[int]()

	set.Load(s, 1, 2, 3, 4, 6, 5)

	if sort.IsSorted(sort.IntSlice(set.ToSlice(s))) {
		t.Fail()
	}

	s.Del(6)

	if !sort.IsSorted(sort.IntSlice(set.ToSlice(s))) {
		t.Fail()
	}

	for i := 5; i > 0; i-- {
		if v, ok := s.Pop(); !ok || v != i {
			t.Fail()
		}
	}
}

func TestUnion(t *testing.T) {
	t.Parallel()

	a := set.Load(set.NewOrdered[int](), 1, 2, 3)
	b := set.Load(set.NewOrdered[int](), 3, 4, 5)

	c := set.Union(a, b)

	if c.Len() != 5 {
		t.Fail()
	}
}

func TestDiff(t *testing.T) {
	t.Parallel()

	a := set.Load(set.NewOrdered[int](), 1, 2, 3)
	b := set.Load(set.NewOrdered[int](), 3, 4, 5)

	c := set.Diff(a, b)

	if c.Len() != 2 {
		t.Fail()
	}

	res := set.ToSlice(c)

	if res[0] != 1 || res[1] != 2 {
		t.Fail()
	}
}

func TestIntersect(t *testing.T) {
	t.Parallel()

	a := set.Load(set.NewOrdered[int](), 1, 2, 3)
	b := set.Load(set.NewOrdered[int](), 3, 4, 5)

	c := set.Intersect(a, b)

	if c.Len() != 1 {
		t.Fail()
	}

	res := set.ToSlice(c)

	if res[0] != 3 {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	a := set.Load(set.NewUnordered[int](), 1, 2, 3, 4, 5, 6)
	b := set.Load(set.NewOrdered[int](), 3, 4, 5)
	c := set.Load(set.NewOrdered[int](), 5, 6, 0)

	if !set.Contains(a, b) {
		t.Fail()
	}

	if !set.Contains(b, a) {
		t.Fail()
	}

	if set.Contains(a, c) {
		t.Fail()
	}
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	a := set.Load(set.NewUnordered[int](), 1, 2, 3, 4, 5, 6)
	b := set.Load(set.NewOrdered[int](), 8, 7, 6)
	c := set.Load(set.NewOrdered[int](), 7, 8, 9)

	if !set.ContainsAny(a, b) {
		t.Fail()
	}

	if set.ContainsAny(a, c) {
		t.Fail()
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()

	a := set.Load(set.NewUnordered[int](), 1, 2, 3)
	b := set.Load(set.NewOrdered[int](), 1, 2, 3)
	c := set.Load(set.NewOrdered[int](), 1, 2, 4)

	if !set.Equal(a, b) {
		t.Fail()
	}

	if set.Equal(a, c) {
		t.Fail()
	}

	// as c is ordered we can be sure, that its pop very last item
	_, _ = c.Pop()

	if set.Equal(a, c) {
		t.Fail()
	}

	c.Add(3)

	if !set.Equal(a, c) {
		t.Fail()
	}
}
