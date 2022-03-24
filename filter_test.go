package streams

import (
	"math/rand"
	"testing"
)

func TestFloatFilter(t *testing.T) {
	seedOnce.Do(seed)
	l := rand.Intn(200)
	if l == 0 {
		l = 200
	}

	testCase := []streamFilterTestCase[float64]{
		{
			input: randFloat64Slice(l),
			filter: Predicate[float64](func(t float64) bool {
				return t/3.0 == 0
			}),
		},
	}

	for _, cs := range testCase {
		var expected []float64
		for _, i := range cs.input {
			if cs.filter(i) {
				expected = append(expected, i)
			}
		}

		result := ToSlice(Filter(FromSlice(cs.input), cs.filter))

		if len(result) != len(expected) {
			t.Logf("result and expected length are not equal, expected %d, got %d\n", len(expected), len(result))
			t.FailNow()
		}

		for i := range result {
			if result[i] != expected[i] {
				t.Logf("expected %v, got %v\n", expected[i], result[i])
				t.FailNow()
			}
		}
	}
}

func TestUintFilter(t *testing.T) {
	seedOnce.Do(seed)
	l := rand.Intn(200)
	if l == 0 {
		l = 200
	}

	testCase := []streamFilterTestCase[uint64]{
		{
			input: randUint64Slice(l),
			filter: Predicate[uint64](func(t uint64) bool {
				return t%2 == 0
			}),
		},
	}

	for _, cs := range testCase {
		var expected []uint64
		for _, i := range cs.input {
			if cs.filter(i) {
				expected = append(expected, i)
			}
		}

		result := ToSlice(Filter(FromSlice(cs.input), cs.filter))

		if len(result) != len(expected) {
			t.Logf("result and expected length are not equal, expected %d, got %d\n", len(expected), len(result))
			t.FailNow()
		}

		for i := range result {
			if result[i] != expected[i] {
				t.Logf("expected %v, got %v\n", expected[i], result[i])
				t.FailNow()
			}
		}
	}
}
