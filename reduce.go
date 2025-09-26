package collection

// Reduce boils down a slice of elements into a single value.
// It applies an accumulator function against an initial value and each element in the slice (from left to right)
// to reduce it to a single value.
func Reduce[T, U any](source []T, initial U, accumulator func(acc U, item T) U) U {
	current := initial
	for _, item := range source {
		current = accumulator(current, item)
	}
	return current
}
