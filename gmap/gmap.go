package gmap

func Keys[K comparable, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func Vals[K comparable, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func MapK[K, K2 comparable, V any](m map[K]V, m2 map[K]K2) map[K2]V {
	ret := make(map[K2]V, len(m))
	for k, v := range m {
		ret[m2[k]] = v
	}
	return ret
}

func MapKF[K, K2 comparable, V any](m map[K]V, f func(K) K2) map[K2]V {
	ret := make(map[K2]V, len(m))
	for k, v := range m {
		ret[f(k)] = v
	}
	return ret
}

func MapV[K comparable, V, V2 any](m map[K]V, v map[K]V2) map[K]V2 {
	ret := make(map[K]V2, len(m))
	for k := range m {
		ret[k] = v[k]
	}
	return ret
}

func MapVF[K comparable, V, V2 any](m map[K]V, f func(V) V2) map[K]V2 {
	ret := make(map[K]V2, len(m))
	for k, v := range m {
		ret[k] = f(v)
	}
	return ret
}

func ForEach[K comparable, V any](m map[K]V, f func(K, V)) {
	for k, v := range m {
		f(k, v)
	}
}

func ForEachE[K comparable, V any](m map[K]V, f func(K, V) error) error {
	for k, v := range m {
		err := f(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
