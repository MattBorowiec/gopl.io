package popcountLoop

import (
	"gopl.io/ch2/popcount"
	"testing"
)

func benchmarkPopCount(u uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(u)
	}
}

func benchmarkPopCountLoop(u uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(u)
	}
}

func BenchmarkPopCount69(b *testing.B)           { benchmarkPopCount(69, b) }
func BenchmarkPopCount420(b *testing.B)          { benchmarkPopCount(420, b) }
func BenchmarkPopCount6969(b *testing.B)         { benchmarkPopCount(6969, b) }
func BenchmarkPopCount4206969(b *testing.B)      { benchmarkPopCount(4206969, b) }
func BenchmarkPopCount694206942069(b *testing.B) { benchmarkPopCount(694206942069, b) }

func BenchmarkPopCountLoop69(b *testing.B)           { benchmarkPopCountLoop(69, b) }
func BenchmarkPopCountLoop420(b *testing.B)          { benchmarkPopCountLoop(420, b) }
func BenchmarkPopCountLoop6969(b *testing.B)         { benchmarkPopCountLoop(6969, b) }
func BenchmarkPopCountLoop4206969(b *testing.B)      { benchmarkPopCountLoop(4206969, b) }
func BenchmarkPopCountLoop694206942069(b *testing.B) { benchmarkPopCountLoop(694206942069, b) }
