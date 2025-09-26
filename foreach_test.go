package collection

import "testing"

func TestForEach(t *testing.T) {
	t.Run("it iterates over each element", func(t *testing.T) {
		source := []int{1, 2, 3}
		var sum int

		action := func(item int) {
			sum += item
		}

		ForEach(source, action)

		expected := 6
		if sum != expected {
			t.Errorf("expected sum to be %d, got %d", expected, sum)
		}
	})

	t.Run("it does nothing on an empty slice", func(t *testing.T) {
		source := []int{}
		var counter int
		action := func(item int) {
			counter++
		}

		ForEach(source, action)

		if counter != 0 {
			t.Errorf("action should not have been called, counter is %d", counter)
		}
	})
}
