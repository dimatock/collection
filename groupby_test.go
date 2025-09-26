package collection

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	t.Run("groups a slice of strings by their length", func(t *testing.T) {
		source := []string{"apple", "banana", "cat", "dog", "elephant"}
		keySelector := func(s string) int {
			return len(s)
		}

		result := GroupBy(source, keySelector)

		expected := map[int][]string{
			3: {"cat", "dog"},
			5: {"apple"},
			6: {"banana"},
			8: {"elephant"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("groups a slice of structs by a property", func(t *testing.T) {
		type Product struct {
			Name     string
			Category string
		}
		source := []Product{
			{Name: "Laptop", Category: "Electronics"},
			{Name: "Book", Category: "Books"},
			{Name: "Mouse", Category: "Electronics"},
		}

		keySelector := func(p Product) string {
			return p.Category
		}

		result := GroupBy(source, keySelector)

		expected := map[string][]Product{
			"Electronics": {{Name: "Laptop", Category: "Electronics"}, {Name: "Mouse", Category: "Electronics"}},
			"Books":       {{Name: "Book", Category: "Books"}},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("returns an empty map for an empty slice", func(t *testing.T) {
		var source []int
		keySelector := func(i int) int { return i }

		result := GroupBy(source, keySelector)

		if len(result) != 0 {
			t.Errorf("expected an empty map, but got %v", result)
		}
	})
}
