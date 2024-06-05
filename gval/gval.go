package gval

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
