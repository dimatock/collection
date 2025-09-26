package collection

// ForEach iterates over a slice and calls the provided action function for each element.
func ForEach[T any](source []T, action func(item T)) {
	for _, item := range source {
		action(item)
	}
}
