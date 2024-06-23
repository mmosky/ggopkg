package gslice

// ForEach iterates over the elements of the slice arr and applies the
// function f to each element.
func ForEach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

// ForEachE iterates over the elements of the slice arr and applies the
// function f to each element. If the function f returns an error for any
// element, ForEachE immediately returns that error. Otherwise, it returns
// nil.
func ForEachE[T any](arr []T, f func(T) error) error {
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

// ForEachE2 iterates over the elements of the slice arr and applies the
// function f to each element and its index. The function f is called with
// the index i and the element v for each element in the slice. If the
// function f returns an error for any element, ForEachE2 immediately
// returns that error. Otherwise, it returns nil.
func ForEachE2[T any](arr []T, f func(int, T) error) error {
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

// FilterE returns a new slice containing all elements of the slice arr for
// which the predicate function pred returns true. If arr is nil, FilterE
// returns nil. If the predicate function pred returns an error for any
// element, FilterE immediately returns that error. Otherwise, it returns nil.
func FilterE[T any](arr []T, pred func(T) (bool, error)) ([]T, error) {
	if arr == nil {
		return nil, nil
	}
	ret := make([]T, 0, len(arr))
	for i, v := range arr {
		b, err := pred(v)
		if err != nil {
			return nil, err
		}
		if b {
			ret = append(ret, arr[i])
		}
	}
	return ret, nil
}

// Filter2 returns two new slices. The first slice contains all elements of
// the slice arr for which the predicate function pred returns true. The
// second slice contains all elements for which the predicate function pred
// returns false. If arr is nil, Filter2 returns two nil slices.
func Filter2[T any](arr []T, pred func(int, T) bool) ([]T, []T) {
	if arr == nil {
		return nil, nil
	}
	ret1 := make([]T, 0, len(arr))
	ret2 := make([]T, 0, len(arr))
	for i, v := range arr {
		if pred(i, v) {
			ret1 = append(ret1, arr[i])
		} else {
			ret2 = append(ret2, arr[i])
		}
	}
	return ret1, ret2
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

// MapE returns a new slice containing the results of applying the function f
// to each element of the slice arr. If arr is nil, MapE returns nil. If the
// function f returns an error for any element, MapE immediately returns that
// error. Otherwise, it returns nil.
func MapE[T1, T2 any](arr []T1, f func(T1) (T2, error)) ([]T2, error) {
	if arr == nil {
		return nil, nil
	}
	ret := make([]T2, len(arr))
	for i, v := range arr {
		v2, err := f(v)
		if err != nil {
			return nil, err
		}
		ret[i] = v2
	}
	return ret, nil
}

// Count returns the number of elements in the slice arr for which the
// predicate function pred returns true.
func Count[T any](arr []T, pred func(T) bool) int {
	count := 0
	for _, v := range arr {
		if pred(v) {
			count++
		}
	}
	return count
}

// In returns true if the slice arr contains the value v, and false otherwise.
func In[T comparable](arr []T, v T) bool {
	for _, e := range arr {
		if e == v {
			return true
		}
	}
	return false
}
