package collection

// Map applies a transformer function to each element of the source slice
// and returns a new slice containing the transformed elements.
func Map[T, U any](source []T, transform func(item T) U) []U {
	if source == nil {
		return nil
	}
	result := make([]U, len(source))
	for i, item := range source {
		result[i] = transform(item)
	}
	return result
}
