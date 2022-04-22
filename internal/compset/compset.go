package compset

type CompSet[T comparable] map[T]struct{}

func (cs CompSet[T]) Has(v T) bool {
	_, ok := cs[v]
	return ok
}

func (cs CompSet[T]) Add(v T) {
	cs[v] = struct{}{}
}

func FromString(s string) CompSet[string] {
	cs := make(CompSet[string])

	for _, r := range s {
		c := string(r)
		cs.Add(c)
	}

	return cs
}

func ToString(cs CompSet[string]) string {
	s := ""

	for k := range cs {
		s += k
	}

	return s
}

func UniqueLetters(s string) string {
	return ToString(FromString(s))
}
