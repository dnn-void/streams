package streams

type Mapper[I, U any] func(I) U

func Map[I, O any](s *Stream[I], f Mapper[I, O]) *Stream[O] {
	if s == nil {
		return nil
	}

	val, ok := s.Value()
	if !ok {
		return nil
	}

	return &Stream[O]{
		Val: f(val),
		Next: func() *Stream[O] {
			next := s.Next()
			if next == nil {
				return nil
			}

			nextVal, ok := next.Value()
			if !ok {
				return nil
			}

			return &Stream[O]{
				Val: f(nextVal),
				Next: func() *Stream[O] {
					return Map(next.Next(), f)
				},
			}
		},
	}
}
