package tuple

type T[T1, T2 any] struct {
	First  T1
	Second T2
}

func (t *T[T1, T2]) Vals() (T1, T2) {
	return t.First, t.Second
}

type T3[T1, T2, T3_ any] struct {
	First  T1
	Second T2
	Third  T3_
}

func (t *T3[T1, T2, T3_]) Vals() (T1, T2, T3_) {
	return t.First, t.Second, t.Third
}
