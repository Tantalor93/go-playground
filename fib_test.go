package benchmark_action_poc

import "testing"

func BenchmarkFib_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}
