package util

import (
	"math"
	"sort"
)

type direction int

const (
	Left direction = iota
	Right
)

func MapValues[K comparable, V any](m map[K]V, f func(V, K) V) map[K]V {
	for k, v := range m {
		m[k] = f(v, k)
	}

	return m
}

func MapSlice[T any, U any](s []T, m func(T, int) U) []U {
	newS := make([]U, len(s))

	for i, v := range s {
		newS[i] = m(v, i)
	}

	return newS
}

func Take[T any](s []T, n uint) []T {
	lim := int(
		math.Min(float64(n), float64(len(s))),
	)
	newS := make([]T, lim)

	for i := 0; i < lim; i++ {
		newS[i] = s[i]
	}

	return newS
}

func Filter[T any](s []T, p func(T) bool) []T {
	newS := make([]T, 0, len(s))

	for _, v := range s {
		if p(v) {
			newS = append(newS, v)
		}
	}

	return newS
}

func IndexOf[T comparable](arr []T, v T) int {
	for i, e := range arr {
		if e == v {
			return i
		}
	}

	return -1
}

func wrap(n int, mod int) int {
	for n < 0 {
		n += mod
	}

	return n % mod
}

func shift(n int, by int, dir direction) int {
	if dir == Left {
		return n - by
	} else if dir == Right {
		return n + by
	}

	return n
}

func ShiftWrap(n int, by int, mod int, dir direction) int {
	shifted := shift(n, by, dir)
	return wrap(shifted, mod)
}

func SplitAfterIndices[T any](s []T, ps ...int) [][]T {
	ps = Filter(ps, func(p int) bool {
		return p <= len(s)
	})
	ss := make([][]T, len(ps)+1)
	prev := 0

	sort.Ints(ps)

	for i, p := range ps {
		ss[i] = s[prev:p]
		prev = p
	}

	ss[len(ps)] = s[prev:]

	return ss
}

func SplitEvery[T any](s []T, i int) [][]T {
	iter := len(s) / i

	if len(s)%i != 0 {
		iter++
	}

	ss := make([][]T, iter)

	for n := 0; n < iter; n++ {
		upper := int(math.Min(float64((n+1)*i), float64(len(s))))
		segment := s[n*i : upper]
		v := make([]T, len(segment))
		copy(v, segment)
		ss[n] = v
	}

	return ss
}

func Concat[T any](a []T, b []T) []T {
	c := make([]T, len(a)+len(b))

	copy(c, a)
	copy(c[len(a):], b)

	return c
}

func Flatten[T any](es [][]T) []T {
	var size int

	for _, e := range es {
		size += len(e)
	}

	e := make([]T, size)
	counter := 0

	for _, v := range es {
		copy(e[counter:counter+len(v)], v)
		counter += len(v)
	}

	return e
}

func Reverse[T any](as []T) []T {
	bs := make([]T, len(as))

	for i, a := range as {
		bs[len(bs)-i-1] = a
	}

	return bs
}

func Every[T any](as []T, f func(a T, i int) bool) bool {
	for i, a := range as {
		if !f(a, i) {
			return false
		}
	}

	return true
}

func AreSlicesEqual[T comparable](as []T, bs []T) bool {
	if len(as) != len(bs) {
		return false
	}

	return Every(as, func(a T, i int) bool {
		return a == bs[i]
	})
}

func Fill[T any](as []T, v T) []T {
	bs := make([]T, len(as), cap(as))

	for i := range as {
		bs[i] = v
	}

	return bs
}

func Last[T any](as []T) T {
	return as[len(as)-1]
}

func FindLast[T any](as []T, f func(a T, i int) bool) (T, int, bool) {
	var zero T

	for i, v := range Reverse(as) {
		if f(v, i) {
			return v, len(as) - i - 1, true
		}
	}

	return zero, 0, false
}
