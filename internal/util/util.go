package util

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
	newS := make([]T, n)

	for i := uint(0); i < n; i++ {
		newS[i] = s[i]
	}

	return newS
}
