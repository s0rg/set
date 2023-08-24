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
func (u Unordered[T]) Add(v T) bool {
	prev := len(u)
	u[v] = stub{}
	return prev != len(u)
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

// Has implements Set interface.
func (u Unordered[T]) Has(v T) (ok bool) {
	_, ok = u[v]

	return
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
	if len(u) < 1 {
		return v, false
	}

	for t := range u {
		v = t

		break
	}

	delete(u, v)
	// u.Del(v)

	return v, true
}
