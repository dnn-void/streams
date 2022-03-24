package streams

func FromSlice[T any](s []T) *Stream[T] {
	if len(s) == 0 {
		return nil
	}

	return &Stream[T]{
		Val: s[0],
		Next: func() *Stream[T] {
			if len(s) == 1 {
				return nil
			}
			return FromSlice(s[1:])
		},
	}
}

func FromChan[T any](c <-chan T) *Stream[T] {
	return &Stream[T]{
		Next: func() *Stream[T] {
			return FromChan(c)
		},
		Getter: func() (T, bool) {
			v, ok := <-c
			return v, ok
		},
	}
}
