package streams

import (
	"math/rand"
	"sync"
	"time"
)

var seedOnce sync.Once

func seed() {
	rand.Seed(time.Now().UnixNano())
}

func randFloat64Slice(l int) []float64 {
	result := make([]float64, l, l)
	for i := range result {
		result[i] = rand.Float64()
	}

	return result
}

func randUint64Slice(l int) []uint64 {
	result := make([]uint64, l, l)
	for i := range result {
		result[i] = rand.Uint64()
	}

	return result
}

func randInt64Slice(l int) []int64 {
	result := make([]int64, l, l)
	for i := range result {
		result[i] = rand.Int63()
	}

	return result
}

type streamFilterTestCase[T any] struct {
	input  []T
	filter Predicate[T]
}
