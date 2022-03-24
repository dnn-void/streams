package streams

type Predicate[T any] func(T) bool

func Filter[T any](s *Stream[T], f Predicate[T]) *Stream[T] {
	if s == nil {
		return nil
	}

	val, ok := s.Value()
	if !ok {
		return nil
	}

	if !f(val) {
		return Filter(s.Next(), f)
	}

	return &Stream[T]{
		Val: val,
		Next: func() *Stream[T] {
			return Filter(s.Next(), f)
		},
	}
}
