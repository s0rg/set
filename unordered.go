package set

type (
	stub struct{}

	// Unordered is a simple and effective un-ordered generic set.
	Unordered[T comparable] map[T]stub
)

// NewUnordered create empty unordered Set for given type.
func NewUnordered[T comparable]() Set[T] {
	return make(Unordered[T])
}

// Add implements Set interface.
func (u Unordered[T]) Add(v T) {
	u[v] = stub{}
}

// Del implements Set interface.
func (u Unordered[T]) Del(v T) {
	delete(u, v)
}

// Clear implements Set interface.
func (u Unordered[T]) Clear() {
	clear(u)
}

// Len implements Set interface.
func (u Unordered[T]) Len() int {
	return len(u)
}

// TryAdd implements Set interface.
func (u Unordered[T]) TryAdd(v T) (ok bool) {
	if u.Has(v) {
		return false
	}

	u.Add(v)

	return true
}

// Has implements Set interface.
func (u Unordered[T]) Has(v T) (ok bool) {
	_, ok = u[v]

	return
}

// ToSlice implements Set interface.
func (u Unordered[T]) ToSlice() (rv []T) {
	rv = make([]T, 0, len(u))

	for k := range u {
		rv = append(rv, k)
	}

	return rv
}

// Iter implements Set interface.
func (u Unordered[T]) Iter(it func(T) bool) {
	for k := range u {
		if !it(k) {
			break
		}
	}
}

// Clone implements Set interface.
func (u Unordered[T]) Clone() (rv Set[T]) {
	rv = make(Unordered[T])

	u.Iter(func(v T) bool {
		rv.Add(v)

		return true
	})

	return rv
}

// Pop implements Set interface.
func (u Unordered[T]) Pop() (v T, ok bool) {
	if u.Len() < 1 {
		return v, false
	}

	u.Iter(func(t T) bool {
		v = t

		return false
	})

	u.Del(v)

	return v, true
}
