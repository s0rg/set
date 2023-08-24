package set_test

import (
	"testing"

	"github.com/s0rg/set"
)

func benchmarkSetMain(b *testing.B, s set.Set[int]) {
	b.Helper()

	b.Run("Add", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = s.Add(n)
		}
	})

	b.Run("Has", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = s.Has(n)
		}
	})

	b.Run("Len", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = s.Len()
		}
	})

	b.Run("Pop", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = s.Pop()
		}
	})
}

func benchmarkSetDel(b *testing.B, s set.Set[int]) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		s.Add(n)
	}

	b.ResetTimer()

	b.Run("Del", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s.Del(n)
		}
	})
}

func BenchmarkSetUnorderedDirect(b *testing.B) {
	s := make(set.Unordered[int])

	b.ResetTimer()

	b.Run("Add", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = s.Add(n)
		}
	})

	b.Run("Has", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = s.Has(n)
		}
	})

	b.Run("Len", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = s.Len()
		}
	})

	b.Run("Pop", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = s.Pop()
		}
	})

	b.StopTimer()

	for n := 0; n < b.N; n++ {
		s.Add(n)
	}

	b.StartTimer()

	b.Run("Del", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s.Del(n)
		}
	})
}

func BenchmarkSetUnorderedIndirect(b *testing.B) {
	s := set.NewUnordered[int]()

	b.ResetTimer()

	benchmarkSetMain(b, s)
	benchmarkSetDel(b, s)
}

func BenchmarkSetOrdered(b *testing.B) {
	s := set.NewOrdered[int]()

	b.ResetTimer()

	benchmarkSetMain(b, s)
	benchmarkSetDel(b, s)
}
