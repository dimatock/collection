package collection

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("finds an element that exists", func(t *testing.T) {
		source := []string{"apple", "banana", "cherry"}
		predicate := func(s string) bool {
			return strings.HasPrefix(s, "b")
		}

		result, found := Find(source, predicate)

		if !found {
			t.Fatal("expected to find an element, but did not")
		}
		if result != "banana" {
			t.Errorf("expected 'banana', got '%s'", result)
		}
	})

	t.Run("returns false when element does not exist", func(t *testing.T) {
		source := []string{"apple", "banana", "cherry"}
		predicate := func(s string) bool {
			return strings.HasPrefix(s, "x")
		}

		_, found := Find(source, predicate)

		if found {
			t.Error("expected not to find an element, but did")
		}
	})

	t.Run("finds the first matching element", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5, 6}
		// Predicate for the first number greater than 3
		predicate := func(i int) bool {
			return i > 3
		}

		result, found := Find(source, predicate)

		if !found {
			t.Fatal("expected to find an element > 3")
		}
		if result != 4 {
			t.Errorf("expected to find 4, but got %d", result)
		}
	})
}
