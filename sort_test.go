package collection

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("sorts a slice of integers", func(t *testing.T) {
		source := []int{5, 2, 8, 1, 9}
		less := func(a, b int) bool {
			return a < b
		}

		result := Sort(source, less)
		expected := []int{1, 2, 5, 8, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected sorted slice %v, got %v", expected, result)
		}
	})

	t.Run("does not modify the original slice", func(t *testing.T) {
		source := []int{3, 1, 2}
		originalSource := make([]int, len(source))
		copy(originalSource, source)

		less := func(a, b int) bool { return a < b }
		Sort(source, less)

		if !reflect.DeepEqual(source, originalSource) {
			t.Errorf("original slice was modified, expected %v, got %v", originalSource, source)
		}
	})

	t.Run("sorts a slice of structs", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		source := []User{{Name: "Bob", Age: 30}, {Name: "Alice", Age: 25}}
		less := func(a, b User) bool {
			return a.Age < b.Age
		}

		result := Sort(source, less)
		expected := []User{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected sorted slice %v, got %v", expected, result)
		}
	})
}
