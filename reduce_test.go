package collection

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("sums up a slice of integers", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5}
		initial := 0
		accumulator := func(acc int, item int) int {
			return acc + item
		}

		result := Reduce(source, initial, accumulator)
		expected := 15

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("concatenates a slice of strings", func(t *testing.T) {
		source := []string{"a", "b", "c"}
		initial := "_"
		accumulator := func(acc string, item string) string {
			return acc + item
		}

		result := Reduce(source, initial, accumulator)
		expected := "_abc"

		if result != expected {
			t.Errorf("expected %s, got %s", expected, result)
		}
	})

	t.Run("returns initial value for an empty slice", func(t *testing.T) {
		source := []int{}
		initial := 100
		accumulator := func(acc int, item int) int {
			return acc + item
		}

		result := Reduce(source, initial, accumulator)

		if result != initial {
			t.Errorf("expected initial value %d, got %d", initial, result)
		}
	})

	t.Run("transforms types during reduction", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		source := []User{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
		initial := 0
		accumulator := func(acc int, item User) int {
			return acc + item.Age
		}

		result := Reduce(source, initial, accumulator)
		expected := 55

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})
}
