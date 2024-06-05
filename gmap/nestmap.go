package gmap

func GetSetNested[K, K2 comparable, V any](m map[K]map[K2]V, k K) map[K2]V {
	if m == nil {
		return nil
	}
	if v, ok := m[k]; ok {
		return v
	}
	m[k] = make(map[K2]V)
	return m[k]
}
