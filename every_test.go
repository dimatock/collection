package collection

import "testing"

func TestEvery(t *testing.T) {
	t.Run("returns true when all elements match", func(t *testing.T) {
		source := []int{2, 4, 6, 8}
		// Check if all numbers are even
		predicate := func(i int) bool { return i%2 == 0 }

		if !Every(source, predicate) {
			t.Error("expected true, got false")
		}
	})

	t.Run("returns false when at least one element does not match", func(t *testing.T) {
		source := []int{2, 4, 5, 8}
		predicate := func(i int) bool { return i%2 == 0 }

		if Every(source, predicate) {
			t.Error("expected false, got true")
		}
	})

	t.Run("returns true for an empty slice", func(t *testing.T) {
		source := []int{}
		predicate := func(i int) bool { return false }

		// The condition is vacuously true for an empty set
		if !Every(source, predicate) {
			t.Error("expected true for empty slice, got false")
		}
	})
}
