package collection

import "testing"

func TestSome(t *testing.T) {
	t.Run("returns true when at least one element matches", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5}
		// Check if there is any even number
		predicate := func(i int) bool { return i%2 == 0 }

		if !Some(source, predicate) {
			t.Error("expected true, got false")
		}
	})

	t.Run("returns false when no elements match", func(t *testing.T) {
		source := []int{1, 3, 5, 7}
		predicate := func(i int) bool { return i%2 == 0 }

		if Some(source, predicate) {
			t.Error("expected false, got true")
		}
	})

	t.Run("returns false for an empty slice", func(t *testing.T) {
		source := []int{}
		predicate := func(i int) bool { return true }

		if Some(source, predicate) {
			t.Error("expected false for empty slice, got true")
		}
	})
}
