package set

// Set is the primary interface provided by the set package.
type Set[T comparable] interface {
	// Add adds item to the set.
	Add(T)
	// Del removes item, no-op if not present.
	Del(T)
	// Has checks if item is already present.
	Has(T) bool
	// TryAdd takes attempt to add item, returns false if it already there.
	TryAdd(T) bool
	// Pop removes and returns an arbitrary item.
	Pop() (v T, ok bool)
	// Len returns current items count.
	Len() int
	// Clone returns shallow copy.
	Clone() Set[T]
	// Iter iterates items.
	Iter(func(T) bool)
	// ToSlice returns set as slice of items.
	ToSlice() []T
	// Clear removes all items.
	Clear()
}

// Load populates given set with values.
func Load[T comparable](s Set[T], v ...T) Set[T] {
	for _, i := range v {
		s.Add(i)
	}

	return s
}

// Union returns new set with elements from both sets.
func Union[T comparable](a, b Set[T]) (rv Set[T]) {
	rv = a.Clone()

	b.Iter(func(v T) bool {
		rv.Add(v)

		return true
	})

	return rv
}

// Diff returns new set with elements from `a` that arent in `b`.
func Diff[T comparable](a, b Set[T]) (rv Set[T]) {
	rv = a.Clone()

	a.Iter(func(v T) bool {
		if b.Has(v) {
			rv.Del(v)
		}

		return true
	})

	return rv
}

// Intersect returns new set with keys from `a` that present in `b`.
func Intersect[T comparable](a, b Set[T]) (rv Set[T]) {
	rv = a.Clone()

	a.Iter(func(v T) bool {
		if !b.Has(v) {
			rv.Del(v)
		}

		return true
	})

	return rv
}

// Contains returns true if smallest of two sets (by length) fully contains inside bigger one,
// if sets equals in length the result is same as comparision.
func Contains[T comparable](a, b Set[T]) (yes bool) {
	if b.Len() > a.Len() {
		a, b = b, a
	}

	b.Iter(func(v T) bool {
		yes = a.Has(v)

		return yes
	})

	return yes
}

// ContainsAny returns true if `a` has at least one element from `b`.
func ContainsAny[T comparable](a, b Set[T]) (yes bool) {
	a.Iter(func(v T) bool {
		yes = b.Has(v)

		return !yes
	})

	return yes
}

// Equal returns true if `a` and `b` has same length and elements.
func Equal[T comparable](a, b Set[T]) (yes bool) {
	return a.Len() == b.Len() && Contains(a, b)
}
