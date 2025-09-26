package collection

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	t.Run("flattens a slice of integer slices", func(t *testing.T) {
		source := [][]int{{1, 2}, {3, 4, 5}, {6}}
		result := Flatten(source)
		expected := []int{1, 2, 3, 4, 5, 6}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("handles slices with empty sub-slices", func(t *testing.T) {
		source := [][]string{{"a"}, {}, {"b", "c"}}
		result := Flatten(source)
		expected := []string{"a", "b", "c"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("handles an empty top-level slice", func(t *testing.T) {
		var source [][]int
		result := Flatten(source)
		if len(result) != 0 {
			t.Errorf("expected an empty slice, but got %v", result)
		}
	})
}
