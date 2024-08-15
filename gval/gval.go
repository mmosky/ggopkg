package gval

import (
	"github.com/mmosky/ggopkg/tuple"
)

func Zero[T any]() T {
	var t T
	return t
}

func Ptr[T any](x T) *T {
	return &x
}

func DeRef[T any](x *T, defaultVal ...T) T {
	if x == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return Zero[T]()
	}
	return *x
}

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func Must2[T, T2 any](x T, y T2, err error) tuple.T[T, T2] {
	if err != nil {
		panic(err)
	}
	return tuple.T[T, T2]{First: x, Second: y}
}

func Must3[T, T2, T3 any](x T, y T2, z T3, err error) tuple.T3[T, T2, T3] {
	if err != nil {
		panic(err)
	}
	return tuple.T3[T, T2, T3]{First: x, Second: y, Third: z}
}
