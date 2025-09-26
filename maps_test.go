package collection

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	t.Run("returns keys from a map of strings to ints", func(t *testing.T) {
		source := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		result := Keys(source)
		expected := []string{"apple", "banana", "cherry"}

		// Sort both slices for stable comparison
		sort.Strings(result)
		sort.Strings(expected)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("returns nil for a nil map", func(t *testing.T) {
		var source map[int]string = nil
		result := Keys(source)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})

	t.Run("returns an empty slice for an empty map", func(t *testing.T) {
		source := map[int]string{}
		result := Keys(source)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})
}

func TestValues(t *testing.T) {
	t.Run("returns values from a map of strings to ints", func(t *testing.T) {
		source := map[string]int{
			"a": 10,
			"b": 20,
			"c": 30,
		}

		result := Values(source)
		expected := []int{10, 20, 30}

		// Sort both slices for stable comparison
		sort.Ints(result)
		sort.Ints(expected)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}

func TestMapValues(t *testing.T) {
	t.Run("transforms the values of a map", func(t *testing.T) {
		source := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		transform := func(v int) string {
			return fmt.Sprintf("v%d", v)
		}

		result := MapValues(source, transform)
		expected := map[string]string{
			"a": "v1",
			"b": "v2",
			"c": "v3",
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
