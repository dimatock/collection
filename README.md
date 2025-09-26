# collection

[![Go Report Card](https://goreportcard.com/badge/github.com/dimatock/collection)](https://goreportcard.com/report/github.com/dimatock/collection)

`collection` is a functional-style utility package for working with slices in Go, written with generics. It provides a simple and powerful API for common collection operations without mutating the source data.

## Philosophy

* **Immutability**: No function in this package modifies the original slice. Every operation returns a new, modified slice.
* **Simplicity**: A clean, minimalistic, and intuitive API.
* **Type-Safety**: Full type safety thanks to Go generics.

## Installation

```bash
go get github.com/dimatock/collection
```

## Testing

To run all tests:

```bash
go test -v ./...
```

To run benchmarks:

```bash
go test -bench=. -benchmem ./...
```

## Performance

Benchmarks show that the overhead of using this package is minimal compared to handwritten `for` loops. For complex operations like `Uniq`, the provided implementation is significantly faster than a naive approach for datasets with moderate to high cardinality.

Key benchmark results on a slice of 1000 elements:

```
goos: darwin
goarch: arm64
pkg: github.com/dimatock/collection

# Map is as fast as a manual for loop
BenchmarkMap/collection.Map-12                   86782      13814 ns/op
BenchmarkMap/for_loop-12                         87528      13880 ns/op

# Filter is as fast as a manual for loop
BenchmarkFilter/collection.Filter-12            775660       1571 ns/op
BenchmarkFilter/for_loop-12                     743002       1665 ns/op

# Optimized Uniq is ~4x faster than a naive O(n*k) approach
BenchmarkUniq/collection.Uniq-12            110916      11244 ns/op
BenchmarkUniq/naive_for_loop-12              13616      84245 ns/op
```

## API & Examples

### `Map`

Transforms each element of a slice into a new value using a transformer function.

**Signature:** `func Map[T, U any](source []T, transform func(item T) U) []U`

```go
source := []int{1, 2, 3}
result := collection.Map(source, func(x int) string {
    return "v" + strconv.Itoa(x)
})
// result: []string{"v1", "v2", "v3"}
```

### `Filter`

Returns a new slice containing only the elements that satisfy the predicate.

**Signature:** `func Filter[T any](source []T, predicate func(item T) bool) []T`

```go
source := []int{1, 2, 3, 4, 5}
result := collection.Filter(source, func(x int) bool {
    return x%2 == 0 // Keep only even numbers
})
// result: []int{2, 4}
```

### `Reduce`

Boils down a slice into a single value using an accumulator function.

**Signature:** `func Reduce[T, U any](source []T, initial U, accumulator func(acc U, item T) U) U`

```go
source := []int{1, 2, 3, 4, 5}
sum := collection.Reduce(source, 0, func(acc int, item int) int {
    return acc + item
})
// sum: 15
```

### `ForEach`

Executes an action for each element of the slice. Returns nothing.

**Signature:** `func ForEach[T any](source []T, action func(item T))`

```go
source := []string{"a", "b"}
collection.ForEach(source, func(item string) {
    fmt.Println(item)
})
// Output:
// a
// b
```

### `Find`

Finds the first element that satisfies the predicate.

**Signature:** `func Find[T any](source []T, predicate func(item T) bool) (T, bool)`

```go
source := []int{1, 2, 3, 4}
even, found := collection.Find(source, func(x int) bool {
    return x%2 == 0
})
// even: 2, found: true
```

### `Some`

Returns `true` if at least one element satisfies the predicate.

**Signature:** `func Some[T any](source []T, predicate func(item T) bool) bool`

```go
source := []int{1, 3, 5, 6}
hasEven := collection.Some(source, func(x int) bool {
    return x%2 == 0
})
// hasEven: true
```

### `Every`

Returns `true` if all elements satisfy the predicate.

**Signature:** `func Every[T any](source []T, predicate func(item T) bool) bool`

```go
source := []int{2, 4, 6}
allEven := collection.Every(source, func(x int) bool {
    return x%2 == 0
})
// allEven: true
```

### `Uniq`

Removes duplicate values from a slice, preserving order.

**Signature:** `func Uniq[T comparable](source []T) []T`

```go
source := []int{1, 2, 2, 3, 1}
unique := collection.Uniq(source)
// unique: []int{1, 2, 3}
```

### `GroupBy`

Groups elements into a `map` where the key is determined by the `keySelector` function.

**Signature:** `func GroupBy[T any, K comparable](source []T, keySelector func(item T) K) map[K][]T`

```go
source := []string{"apple", "banana", "cat", "dog"}
grouped := collection.GroupBy(source, func(s string) int {
    return len(s)
})
// grouped: map[int][]string{
//     3: {"cat", "dog"},
//     5: {"apple"},
//     6: {"banana"},
// }
```

### `Flatten`

Flattens a slice of slices into a single slice.

**Signature:** `func Flatten[T any](source [][]T) []T`

```go
source := [][]int{{1, 2}, {3, 4}}
flat := collection.Flatten(source)
// flat: []int{1, 2, 3, 4}
```

### `Sort`

Returns a new, sorted copy of the slice.

**Signature:** `func Sort[T any](source []T, less func(a, b T) bool) []T`

```go
source := []int{3, 1, 2}
sorted := collection.Sort(source, func(a, b int) bool {
    return a < b
})
// sorted: []int{1, 2, 3}
// source: []int{3, 1, 2} (is not modified)
```

## Map Functions

### `Keys`

Returns a slice of the map's keys.

**Signature:** `func Keys[K comparable, V any](source map[K]V) []K`

```go
source := map[string]int{"a": 1, "b": 2}
keys := collection.Keys(source)
// keys: []string{"a", "b"} (order is not guaranteed)
```

### `Values`

Returns a slice of the map's values.

**Signature:** `func Values[K comparable, V any](source map[K]V) []V`

```go
source := map[string]int{"a": 1, "b": 2}
values := collection.Values(source)
// values: []int{1, 2} (order is not guaranteed)
```

### `MapValues`

Creates a new map with the same keys but with transformed values.

**Signature:** `func MapValues[K comparable, V, U any](source map[K]V, transform func(val V) U) map[K]U`

```go
source := map[string]int{"a": 1, "b": 2}
mapped := collection.MapValues(source, func(v int) string {
    return fmt.Sprintf("val_%d", v)
})
// mapped: map[string]string{"a": "val_1", "b": "val_2"}
```

