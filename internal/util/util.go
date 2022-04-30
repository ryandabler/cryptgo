package util

import "math"

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

func MapSlice[T any, U any](s []T, m func(T) U) []U {
	newS := make([]U, len(s))

	for i, v := range s {
		newS[i] = m(v)
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
