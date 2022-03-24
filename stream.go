package streams

type Stream[T any] struct {
	Val    T
	Next   func() *Stream[T]
	Getter func() (T, bool)
}

func (s *Stream[T]) Value() (T, bool) {
	if s.Getter != nil {
		return s.Getter()
	}
	return s.Val, true
}
