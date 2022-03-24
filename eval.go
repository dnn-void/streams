package streams

type Reducer[T, R any] func(acc R, cur T) R

func Reduce[T, R any](s *Stream[T], acc R, f Reducer[T, R]) R {
	if s == nil {
		return acc
	}

	IterValues(s, func(t T) {
		acc = f(acc, t)
	})

	return acc
}

func ToSlice[T any](s *Stream[T]) []T {
	var result []T

	IterValues(s, func(t T) {
		result = append(result, t)
	})

	return result
}

func ToChan[T any](s *Stream[T]) <-chan T {
	ch := make(chan T)
	go func(c chan T) {
		IterValues(s, func(t T) {
			ch <- t
		})
		close(c)
	}(ch)
	return ch
}
