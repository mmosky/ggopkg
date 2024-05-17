package gslice

func ForEach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func TryForEach[T any](arr []T, f func(T) error) error {
	for _, v := range arr {
		err := f(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func ForEach2[T any](arr []T, f func(int, T)) {
	for i, v := range arr {
		f(i, v)
	}
}

func TryForEach2[T any](arr []T, f func(int, T) error) error {
	for i, v := range arr {
		err := f(i, v)
		if err != nil {
			return err
		}
	}
	return nil
}

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
