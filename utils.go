package streams

// Iter iterates through stream objects & calls f on each stream
func Iter[T any](sIn *Stream[T], f func(s *Stream[T])) {
	for {
		if sIn == nil {
			break
		}

		f(sIn)

		sIn = sIn.Next()
	}
}

// IterValues iterates through stream values
func IterValues[T any](sIn *Stream[T], f func(T)) {
	Iter(sIn, func(s *Stream[T]) {
		f(s.Val)
	})
}
