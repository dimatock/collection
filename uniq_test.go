package collection

import (
	"reflect"
	"testing"
)

func TestUniq(t *testing.T) {
	t.Run("removes duplicates from a slice of integers", func(t *testing.T) {
		source := []int{1, 2, 2, 3, 1, 4, 5, 5}
		result := Uniq(source)
		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("preserves order", func(t *testing.T) {
		source := []string{"c", "a", "b", "c", "a"}
		result := Uniq(source)
		expected := []string{"c", "a", "b"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("handles a slice with no duplicates", func(t *testing.T) {
		source := []int{1, 2, 3}
		result := Uniq(source)
		expected := []int{1, 2, 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("handles an empty slice", func(t *testing.T) {
		source := []int{}
		result := Uniq(source)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})
}

var benchmarkUniqResult []int

func BenchmarkUniq(b *testing.B) {
	source := make([]int, 1000)
	for i := 0; i < len(source); i++ {
		source[i] = i % 500 // Create a lot of duplicates, but more unique values
	}

	b.Run("collection.Uniq", func(b *testing.B) {
		var r []int
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r = Uniq(source)
		}
		benchmarkUniqResult = r
	})
}
