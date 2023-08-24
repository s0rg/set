package set

import "slices"

type ordered[T comparable] struct {
	set   Set[T]
	order []T
}

// NewOrdered create empty ordered set for given type.
func NewOrdered[T comparable]() Set[T] {
	return &ordered[T]{
		set:   NewUnordered[T](),
		order: []T{},
	}
}

func (o *ordered[T]) Add(v T) {
	prev := o.set.Len()

	o.set.Add(v)

	if o.set.Len() != prev {
		o.order = append(o.order, v)
	}
}

func (o *ordered[T]) Del(v T) {
	prev := o.set.Len()
	o.set.Del(v)

	if o.set.Len() != prev {
		idx := slices.Index(o.order, v)

		o.order = slices.Clip(slices.Delete(o.order, idx, idx+1))
	}
}

func (o *ordered[T]) Clear() {
	o.set.Clear()
	clear(o.order)
}

func (o *ordered[T]) Len() int {
	return o.set.Len()
}

func (o *ordered[T]) TryAdd(v T) (ok bool) {
	if o.Has(v) {
		return false
	}

	o.Add(v)

	return true
}

func (o *ordered[T]) Has(v T) (ok bool) {
	return o.set.Has(v)
}

func (o *ordered[T]) ToSlice() (rv []T) {
	return o.order
}

func (o *ordered[T]) Iter(it func(T) bool) {
	for _, v := range o.order {
		if !it(v) {
			break
		}
	}
}

func (o *ordered[T]) Clone() (rv Set[T]) {
	rv = NewOrdered[T]()

	o.Iter(func(v T) bool {
		rv.Add(v)

		return true
	})

	return rv
}

func (o *ordered[T]) Pop() (v T, ok bool) {
	ln := o.Len()

	if ln < 1 {
		return v, false
	}

	v = o.order[ln-1]

	o.Del(v)

	return v, true
}
