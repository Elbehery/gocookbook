package set

type Set struct {
	elems map[any]struct{}
}

func NewSet() *Set {
	return &Set{
		elems: map[any]struct{}{},
	}
}

func newSet(elems map[any]struct{}) *Set {
	return &Set{elems: elems}
}

func (s *Set) Add(e any) {
	s.elems[e] = struct{}{}
}

func (s *Set) Remove(e any) {
	delete(s.elems, e)
}

func (s *Set) Size() int {
	return len(s.elems)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) List() []any {
	var list []any

	for k := range s.elems {
		list = append(list, k)
	}
	return list
}

func (s *Set) Has(e any) bool {
	_, ok := s.elems[e]
	return ok
}

func (s *Set) Copy() Set {
	u := NewSet()
	for k := range s.elems {
		u.Add(k)
	}
	return *u
}

func (s *Set) Difference(u Set) Set {
	res := s.Copy()

	for k := range res.elems {
		if u.Has(k) {
			res.Remove(k)
		}
	}
	return res
}

func (s *Set) IsSubset(t Set) bool {
	for k := range s.elems {
		if !t.Has(k) {
			return false
		}
	}
	return true
}

func Union(sets ...Set) Set {
	res := sets[0].Copy()

	for _, s := range sets[1:] {
		for k := range s.elems {
			res.Add(k)
		}
	}

	return res
}

func Intersect(sets ...Set) Set {
	u := sets[0].Copy()

	for k := range u.elems {
		for _, s := range sets[1:] {
			if !s.Has(k) {
				u.Remove(k)
			}
		}
	}
	return u
}
