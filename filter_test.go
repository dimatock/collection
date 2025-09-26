package collection

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("filters a slice of integers", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5, 6}
		// Predicate to keep only even numbers
		isEven := func(x int) bool {
			return x%2 == 0
		}

		result := Filter(source, isEven)
		expected := []int{2, 4, 6}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("returns an empty slice when no elements match", func(t *testing.T) {
		source := []int{1, 3, 5}
		isEven := func(x int) bool { return x%2 == 0 }

		result := Filter(source, isEven)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("handles an empty slice", func(t *testing.T) {
		source := []int{}
		predicate := func(x int) bool { return true }

		result := Filter(source, predicate)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("handles a nil slice", func(t *testing.T) {
		var source []int = nil
		predicate := func(x int) bool { return true }

		result := Filter(source, predicate)
		if result != nil {
			t.Errorf("expected nil slice, got %v", result)
		}
	})
}

var benchmarkFilterResult []int

func BenchmarkFilter(b *testing.B) {
	source := make([]int, 1000)
	for i := 0; i < len(source); i++ {
		source[i] = i
	}
	predicate := func(x int) bool {
		return x%2 == 0
	}

	b.Run("collection.Filter", func(b *testing.B) {
		var r []int
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r = Filter(source, predicate)
		}
		benchmarkFilterResult = r
	})

	b.Run("for loop", func(b *testing.B) {
		var r []int
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := make([]int, 0)
			for _, item := range source {
				if predicate(item) {
					result = append(result, item)
				}
			}
			r = result
		}
		benchmarkFilterResult = r
	})
}
