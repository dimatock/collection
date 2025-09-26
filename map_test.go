package collection

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("transforms slice of ints to strings", func(t *testing.T) {
		source := []int{1, 2, 3}
		transform := func(x int) string {
			return "val_" + strconv.Itoa(x)
		}

		result := Map(source, transform)
		expected := []string{"val_1", "val_2", "val_3"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("handles empty slice", func(t *testing.T) {
		source := []int{}
		transform := func(x int) int { return x * 2 }

		result := Map(source, transform)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("handles nil slice", func(t *testing.T) {
		var source []int = nil
		transform := func(x int) int { return x * 2 }

		result := Map(source, transform)
		if result != nil {
			t.Errorf("expected nil slice, got %v", result)
		}
	})
}

var benchmarkMapResult []string

func BenchmarkMap(b *testing.B) {
	source := make([]int, 1000)
	for i := 0; i < len(source); i++ {
		source[i] = i
	}
	transform := strconv.Itoa

	b.Run("collection.Map", func(b *testing.B) {
		var r []string
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r = Map(source, transform)
		}
		benchmarkMapResult = r
	})

	b.Run("for loop", func(b *testing.B) {
		var r []string
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := make([]string, len(source))
			for j, item := range source {
				result[j] = transform(item)
			}
			r = result
		}
		benchmarkMapResult = r
	})
}
