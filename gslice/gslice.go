package gslice

// ForEach iterates over the elements of the slice arr and applies the
// function f to each element.
func ForEach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

// TryForEach iterates over the elements of the slice arr and applies the
// function f to each element. If the function f returns an error for any
// element, TryForEach immediately returns that error. Otherwise, it returns
// nil.
func TryForEach[T any](arr []T, f func(T) error) error {
	for _, v := range arr {
		err := f(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// ForEach2 iterates over the elements of the slice arr and applies the
// function f to each element and its index. The function f is called with
// the index i and the element v for each element in the slice.
func ForEach2[T any](arr []T, f func(int, T)) {
	for i, v := range arr {
		f(i, v)
	}
}

// TryForEach2 iterates over the elements of the slice arr and applies the
// function f to each element and its index. The function f is called with
// the index i and the element v for each element in the slice. If the
// function f returns an error for any element, TryForEach2 immediately
// returns that error. Otherwise, it returns nil.
func TryForEach2[T any](arr []T, f func(int, T) error) error {
	for i, v := range arr {
		err := f(i, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Filter returns a new slice containing all elements of the slice arr for
// which the predicate function pred returns true. If arr is nil, Filter
// returns nil.
func Filter[T any](arr []T, pred func(T) bool) []T {
	if arr == nil {
		return nil
	}
	ret := make([]T, 0, len(arr))
	for i, v := range arr {
		if pred(v) {
			ret = append(ret, arr[i])
		}
	}
	return ret
}

// Map returns a new slice containing the results of applying the function f
// to each element of the slice arr. If arr is nil, Map returns nil.
func Map[T1, T2 any](arr []T1, f func(T1) T2) []T2 {
	if arr == nil {
		return nil
	}
	ret := make([]T2, len(arr))
	for i, v := range arr {
		ret[i] = f(v)
	}
	return ret
}
