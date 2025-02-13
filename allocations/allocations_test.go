package main

import "testing"

// write a benchmark that calls the example(i int) function with i = 1000
// and run it with go test -bench=.
// what is the output?
// (hint: you can use the testing package's Benchmark function)

func BenchmarkExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		example(1)
	}
}

func BenchmarkExample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main1()
	}
}

func BenchmarkExample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main2()
	}
}

func BenchmarkExample3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main3()
	}
}

// Results:
// If I call the example function with i = 1000, the allocs/op (avg allocations in the heap per operation) is 12
// If I call the example function with i = 1, the allocs/op is 1
// This is likely because of the resizing of the array.
// With Big O, this value would be amortized, so space would be O(n) and time would be O(n). Because it doesn't directly increase with scale. However,
// it DOES increase, even if not proportionate to the size of the input. So it would be best to allocate memory on the heap a single time.

// Comparison between tests:
// Notice that the CPU is way higher in example 2 than in example 1 or 3. The memory is the same, but the CPU overhead is more
// BenchmarkExample-10             78495499                15.16 ns/op            8 B/op          1 allocs/op
// BenchmarkExample1-10            602325729                2.009 ns/op           0 B/op          0 allocs/op
// BenchmarkExample2-10            134173951                8.954 ns/op           8 B/op          1 allocs/op
// BenchmarkExample3-10            590713974                2.021 ns/op           0 B/op          0 allocs/op
