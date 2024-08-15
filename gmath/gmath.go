package gmath

import (
	"github.com/mmosky/ggopkg/constraint"
	"github.com/mmosky/ggopkg/gval"
)

func Abs[T constraint.Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the maximum value in the slice x. If the slice is empty, Max
// returns the zero value of T.
func Max[T constraint.Ordered](x ...T) T {
	if len(x) == 0 {
		return gval.Zero[T]()
	}
	max := x[0]
	for _, v := range x[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum value in the slice x. If the slice is empty, Min
// returns the zero value of T.
func Min[T constraint.Ordered](x ...T) T {
	if len(x) == 0 {
		return gval.Zero[T]()
	}
	min := x[0]
	for _, v := range x[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Clamp[T constraint.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func AtLeast[T constraint.Ordered](x, min T) T {
	if x < min {
		return min
	}
	return x
}

func AtMost[T constraint.Ordered](x, max T) T {
	if x > max {
		return max
	}
	return x
}

func Sum[T constraint.Number](x ...T) T {
	var sum T
	for _, v := range x {
		sum += v
	}
	return sum
}

// Avg returns the average of the values in the slice x. If the slice is empty,
// Avg panics.
func Avg[T constraint.Number](x ...T) T {
	return Sum(x...) / T(len(x))
}
